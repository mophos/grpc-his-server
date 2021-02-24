package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/siteslave/grpc-demo/database"
	proto "github.com/siteslave/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedEmrServiceServer
	proto.UnimplementedDoctorServiceServer
}

type Person struct {
	GwRecordId string `gorm:"primary_key" json: "gw_record_id"`
	GwHospcode string `json:"gw_hospcode"`
	PatientHn  string `json:"patient_hn"`
	Cid        string `json:"cid"`
	Pname      string `json:"pname"`
	Fname      string `json:"fname"`
	Lname      string `json:"lname"`
	Birthdate  string `json:"birthdate"`
}

type Service struct {
	GwRecordId string `gorm:"primary_key" json: "gw_record_id"`
	GwHospcode string `json:"gw_hospcode"`
	Hn         string `json:"hn"`
	Vn         string `json:"vn"`
	Vstdate    string `json:"vstdate"`
	Vsttime    string `json:"vsttime"`
	Pttype     string `json:"pttype"`
	Pttypeno   string `json:"pttypeno"`
	Spclty     string `json:"spclty"`
	Hospname   string `json:"hospname"`
}

type Doctor struct {
	GwRecordId string `gorm:"primary_key" json: "gw_record_id"`
	Name       string `json:"name"`
	LicenseNo  string `json:"licenseno"`
	GwHospcode string `json:"gw_hospcode"`
}

type Screening struct {
	GwRecordId string `gorm:"primary_key" json: "gw_record_id"`
	Hn         string `json:"hn"`
	Vn         string `json:"vn"`
	Vstdate    string `json:"vstdate"`
	Vsttime    string `json:"vsttime"`
	Hospname   string `json:"hospname"`
	GwHospcode string `json:"gw_hospcode"`
	Diagename  string `json:"diagename"`
	Diagtname  string `json:"diagtname"`
}

func initDatabase() {
	var err error
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")

	database.DBConn, err = gorm.Open("mysql", user+":"+pass+"@tcp(127.0.0.1:3309)/hosxpv3?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}

func main() {

	initDatabase()
	defer database.DBConn.Close()

	lis, err := net.Listen("tcp", ":4041")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterEmrServiceServer(srv, &server{})
	proto.RegisterDoctorServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(lis); e != nil {
		panic(err)
	}

}

func (s *server) PatientInfo(_ context.Context, request *proto.RequestCid) (*proto.InfoResponse, error) {
	cid := request.GetCid()
	// fmt.Print(cid)

	db := database.DBConn

	db.SingularTable(true)

	person := []*proto.InfoResponse_Info{}

	res := db.Raw(`
	select gw_record_id,gw_hospcode,patient_hn,cid,pname,fname,lname,birthdate from hosxpv3_person
	where cid=?`, cid).Scan(&person)

	// res := db.Table("hosxpv3_person").Where(&Person{Cid: cid}).Find(&person)

	if res.Error != nil {
		fmt.Println(res.Error)
	}

	return &proto.InfoResponse{
		Results: person,
	}, nil

}

func (s *server) DoctorList(_ context.Context, request *proto.RequestHospcode) (*proto.DoctorResponse, error) {
	hospcode := request.GetHospcode()
	// fmt.Print(cid)

	db := database.DBConn

	db.SingularTable(true)

	doctor := []*proto.DoctorResponse_Doctor{}

	// res := db.Table("doctor").Where(&Doctor{GwHospcode: hospcode}).Find(&doctor)
	res := db.Raw(`
	select * from doctor
	where gw_hospcode=?`, hospcode).Scan(&doctor)

	if res.Error != nil {
		fmt.Println(res.Error)
	}

	return &proto.DoctorResponse{
		Results: doctor,
	}, nil

}

func (s *server) GetServices(_ context.Context, request *proto.RequestCid) (*proto.ServiceResponse, error) {
	cid := request.GetCid()
	// fmt.Print(cid)

	db := database.DBConn

	db.SingularTable(true)

	services := []*proto.ServiceResponse_Service{}

	res := db.Raw(`
	select 
	o.gw_record_id, o.hn, o.hospcode, 
	o.vn, o.vstdate, o.vsttime, 
	o.pttype, o.pttypeno, o.spclty, 
	h.hospname 
	from opd_visit as o
	inner join hoscpperson as p on p.patient_hn=o.hn and p.hospcode=o.hospcode
	inner join b_hospitals as h on h.hospcode=o.hospcode
	where p.cid=?
	and LENGTH(o.vstdate) > 0
	order by o.vstdate, o.vsttime desc`, cid).Scan(&services)

	if res.Error != nil {
		fmt.Println(res.Error)
	}

	return &proto.ServiceResponse{
		Services: services,
	}, nil

}

func (s *server) GetScreening(_ context.Context, request *proto.RequestPatient) (*proto.ScreeningResponse, error) {
	// hn := request.GetHn()
	// vn := request.GetVn()
	hospcode := request.GetHospcode()
	// fmt.Print(cid)

	db := database.DBConn

	db.SingularTable(true)

	screenings := []*proto.ScreeningResponse_Screening{}

	res := db.Raw(`
	SELECT
		d.gw_record_id,
		d.hn,
		d.vn,
		d.icd10,
		d.vstdate,
		d.vsttime,
		h.hospname,
		d.gw_hospcode,
		i.diagename,
		i.diagtname 
	FROM
		hosxpv3.hosxpv3_ovstdiag d
		LEFT JOIN MASTER.b_hospitals AS h ON h.hospcode = d.gw_hospcode
		LEFT JOIN MASTER.icd10 AS i ON i.diagcode = d.icd10
		WHERE d.gw_hospcode=?`, hospcode).Scan(&screenings)
	fmt.Print(hospcode)
	// WHERE d.hn=? and d.vn=? and d.gw_hospcode=?`, hn, vn, hospcode).Scan(&services)

	if res.Error != nil {
		fmt.Println(res.Error)
	}

	return &proto.ScreeningResponse{
		Screenings: screenings,
	}, nil

}
