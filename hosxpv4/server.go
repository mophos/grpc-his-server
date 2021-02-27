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
	proto.UnimplementedMasterServiceServer
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

	err := godotenv.Load("conf.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	initDatabase()
	defer database.DBConn.Close()

	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterEmrServiceServer(srv, &server{})
	proto.RegisterMasterServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(lis); e != nil {
		panic(err)
	}

}

func (s *server) PatientInfo(_ context.Context, request *proto.RequestCid) (*proto.InfoResponse, error) {
	cid := request.GetCid()
	db := database.DBConn
	db.SingularTable(true)
	data := []*proto.InfoResponse_Info{}
	res := db.Raw(`
	select gw_record_id,gw_hospcode,patient_hn,cid,pname,fname,lname,birthdate from hosxpv4_person
	where cid=?`, cid).Scan(&data)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return &proto.InfoResponse{
		Results: data,
	}, nil
}

func (s *server) DoctorList(_ context.Context, request *proto.RequestHospcode) (*proto.DoctorResponse, error) {
	hospcode := request.GetHospcode()
	db := database.DBConn
	db.SingularTable(true)
	data := []*proto.DoctorResponse_Doctor{}
	res := db.Raw(`
	select gw_record_id,gw_hospcode,name,licenseno as license_no,cid from hosxpv4_doctor
	where gw_hospcode=?`, hospcode).Scan(&data)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return &proto.DoctorResponse{
		Results: data,
	}, nil
}

func (s *server) GetServices(_ context.Context, request *proto.RequestCid) (*proto.ServiceResponse, error) {
	cid := request.GetCid()
	db := database.DBConn
	db.SingularTable(true)
	data := []*proto.ServiceResponse_Service{}
	res := db.Raw(`
	select 
	o.gw_record_id, o.hn, o.gw_hospcode, 
	o.vn, o.vstdate, o.vsttime, 
	o.pttype, o.pttypeno, o.spclty, 
	h.hospname,p.cid 
	from hosxpv4.hosxpv4_ovst as o
	inner  join hosxpv4.hosxpv4_person as p on p.patient_hn=o.hn 
	left  join master.b_hospitals as h on h.hospcode=o.gw_hospcode
	where p.cid=?
	and LENGTH(o.vstdate) > 0
	order by o.vstdate, o.vsttime desc`, cid).Scan(&data)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return &proto.ServiceResponse{
		Results: data,
	}, nil
}

func (s *server) GetScreening(_ context.Context, request *proto.RequestPatient) (*proto.ScreeningResponse, error) {
	hn := request.GetHn()
	vn := request.GetVn()
	hospcode := request.GetHospcode()
	db := database.DBConn
	db.SingularTable(true)
	data := []*proto.ScreeningResponse_Screening{}
	res := db.Raw(`
	SELECT
		o.gw_record_id,
		o.gw_hospcode,
		o.hn,
		o.vn,
		o.vstdate,
		o.vsttime,
		h.hospname 
	FROM
		hosxpv4_ovst AS o
		JOIN MASTER.b_hospitals AS h ON h.hospcode = o.gw_hospcode
	WHERE o.gw_hospcode=? and o.hn=? and o.vn=?`, hospcode, hn, vn).Scan(&data)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return &proto.ScreeningResponse{
		Results: data,
	}, nil
}

func (s *server) GetDiagnosis(_ context.Context, request *proto.RequestPatient) (*proto.DiagnosisResponse, error) {
	hn := request.GetHn()
	vn := request.GetVn()
	hospcode := request.GetHospcode()
	db := database.DBConn
	db.SingularTable(true)
	data := []*proto.DiagnosisResponse_Diagnosis{}
	res := db.Raw(`
	SELECT
		d.gw_record_id,
		d.gw_hospcode,
		d.hn,
		d.vn,
		d.icd10,
		d.vstdate,
		d.vsttime,
		h.hospname,
		i.diagename,
		i.diagtname 
	FROM
		hosxpv4.hosxpv4_ovstdiag d
		JOIN MASTER.b_hospitals AS h ON h.hospcode = d.gw_hospcode
		LEFT JOIN MASTER.icd10 AS i ON i.diagcode = d.icd10
		WHERE d.gw_hospcode=? and d.hn=? and d.vn=?`, hospcode, hn, vn).Scan(&data)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return &proto.DiagnosisResponse{
		Results: data,
	}, nil
}

func (s *server) ClinicList(_ context.Context, request *proto.RequestHospcode) (*proto.ClinicResponse, error) {
	hospcode := request.GetHospcode()
	fmt.Print(hospcode)
	db := database.DBConn
	db.SingularTable(true)
	data := []*proto.ClinicResponse_Clinic{}
	res := db.Raw(`
	SELECT gw_record_id,gw_hospcode,depcode as clinic_code,department as clinic_name from hosxpv4_kskdepartment WHERE gw_hospcode=?`, hospcode).Scan(&data)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return &proto.ClinicResponse{
		Results: data,
	}, nil
}

func (s *server) GetProcedure(_ context.Context, request *proto.RequestPatient) (*proto.ProcedureResponse, error) {
	hn := request.GetHn()
	vn := request.GetVn()
	hospcode := request.GetHospcode()
	db := database.DBConn
	db.SingularTable(true)
	data := []*proto.ProcedureResponse_Procedure{}
	res := db.Raw(`
	SELECT
		d.gw_record_id,
		d.gw_hospcode,
		d.hn,
		d.vn,
		d.icd10,
		d.vstdate,
		d.vsttime,
		h.hospname,
		i.diagename,
		i.diagtname 
	FROM
		hosxpv4.hosxpv4_ovstdiag d
		JOIN MASTER.b_hospitals AS h ON h.hospcode = d.gw_hospcode
		LEFT JOIN MASTER.icd10 AS i ON i.diagcode = d.icd10
		LEFT JOIN hosxpv4.hosxpv4_ovst AS o ON o.gw_hospcode = d.gw_hospcode 
		AND o.hn = d.hn 
		AND o.vn = d.vn 
	WHERE
		LEFT ( i.diagcode, 1 ) IN ( 1, 2, 3, 4, 5, 6, 7, 8, 9, 0 )
		AND d.gw_hospcode=? and d.hn=? and d.vn=?`, hospcode, hn, vn).Scan(&data)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return &proto.ProcedureResponse{
		Results: data,
	}, nil
}

func (s *server) GetLab(_ context.Context, request *proto.RequestPatient) (*proto.LabResponse, error) {
	hn := request.GetHn()
	vn := request.GetVn()
	hospcode := request.GetHospcode()
	db := database.DBConn
	db.SingularTable(true)
	data := []*proto.LabResponse_Lab{}
	res := db.Raw(`
	SELECT
		lh.gw_record_id,
		lh.gw_hospcode,
		lh.hn,
		lh.vn,
		lo.lab_items_code,
		lo.lab_items_name_ref,
		lo.lab_items_normal_value_ref,
		h.hospname,
		li.lab_items_name,
		li.lab_items_unit,
		li.lab_items_normal_value
	FROM
		hosxpv4.hosxpv4_lab_head AS lh
		LEFT JOIN hosxpv4.hosxpv4_lab_order AS lo ON lo.lab_order_number = lh.lab_order_number
		LEFT JOIN MASTER.b_hospitals AS h ON h.hospcode = lh.gw_hospcode
		LEFT JOIN MASTER.b_lab_items AS li ON li.gw_hospcode = lh.gw_hospcode 
		AND li.lab_items_code = lo.lab_items_code 
	WHERE
		lh.gw_hospcode=? and lh.hn=? and lh.vn=?`, hospcode, hn, vn).Scan(&data)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return &proto.LabResponse{
		Results: data,
	}, nil
}

func (s *server) GetVaccine(_ context.Context, request *proto.RequestPatient) (*proto.VaccineResponse, error) {
	hn := request.GetHn()
	vn := request.GetVn()
	hospcode := request.GetHospcode()
	db := database.DBConn
	db.SingularTable(true)
	data := []*proto.VaccineResponse_Vaccine{}
	res := db.Raw(`
	SELECT
		ov.gw_record_id,
		ov.gw_hospcode,
		h.hospname,
		o.hn,
		o.vn,
		ov.vaccine_note,
		ov.abnormal_note
	FROM
		hosxpv4.ovst_vaccine AS ov
		LEFT JOIN hosxpv4.hosxpv4_ovst AS o ON o.gw_hospcode = ov.gw_hospcode 
		AND o.vn = ov.vn
		LEFT JOIN MASTER.b_hospitals AS h ON h.hospcode = ov.gw_hospcode
		LEFT JOIN hosxpv4.person_epi_vaccine AS pev ON pev.gw_hospcode = ov.gw_hospcode 
		AND pev.person_epi_vaccine_id = ov.person_vaccine_id
	WHERE
		o.gw_hospcode=? and o.hn=? and o.vn=?`, hospcode, hn, vn).Scan(&data)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return &proto.VaccineResponse{
		Results: data,
	}, nil
}

func (s *server) GetDrug(_ context.Context, request *proto.RequestPatient) (*proto.DrugResponse, error) {
	hn := request.GetHn()
	vn := request.GetVn()
	hospcode := request.GetHospcode()
	db := database.DBConn
	db.SingularTable(true)
	data := []*proto.DrugResponse_Drug{}
	res := db.Raw(`
	SELECT
		op.gw_record_id,
		op.gw_hospcode,
		op.hn,
		op.vn,
		op.icode,
		op.qty,
		op.unitprice,
		op.vstdate,
		op.vsttime,
		h.hospname
	FROM
		hosxpv4.hosxpv4_opitemrece AS op
		LEFT JOIN MASTER.b_hospitals AS h ON h.hospcode = op.gw_hospcode
	WHERE
		op.gw_hospcode=? and op.hn=? and op.vn=?`, hospcode, hn, vn).Scan(&data)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return &proto.DrugResponse{
		Results: data,
	}, nil
}
