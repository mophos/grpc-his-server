package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/siteslave/grpc-demo/database"
	proto "github.com/siteslave/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedEmrServiceServer
	proto.UnimplementedDoctorServiceServer
}

func initDatabase() {
	var err error
	db_user := os.Getenv("DB_USERNAME")
	db_pass := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	// siteTitle := os.Getenv("SITE_TITLE")
	// dbHost := os.Getenv("DB_HOST")
	database.DBConn, err = gorm.Open("mysql", db_user+":"+db_pass+"@tcp("+db_host+":"+db_port+")/"+db_name+"?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	initDatabase()
	defer database.DBConn.Close()

	lis, err := net.Listen("tcp", ":4042")
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
	db := database.DBConn
	db.SingularTable(true)
	person := []*proto.InfoResponse_Info{}
	res := db.Raw(`
	select gw_record_id,gw_hospcode,patient_hn,cid,pname,fname,lname,birthdate from hosxpv3_person
	where cid=?`, cid).Scan(&person)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return &proto.InfoResponse{
		Results: person,
	}, nil
}

func (s *server) DoctorList(_ context.Context, request *proto.RequestHospcode) (*proto.DoctorResponse, error) {
	hospcode := request.GetHospcode()
	db := database.DBConn
	db.SingularTable(true)
	doctor := []*proto.DoctorResponse_Doctor{}
	res := db.Raw(`
	select gw_record_id,gw_hospcode,name,licenseno as license_no,cid from hosxpv3_doctor
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
	hospcode := request.GetHospcode()
	fmt.Print(hospcode)
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
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return &proto.ScreeningResponse{
		Screenings: screenings,
	}, nil
}
