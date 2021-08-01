package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/jung-kurt/gofpdf"
)

const ipfsURL = "ipfs:5001"

// const ipfsURL = "localhost:5001"

const (
	Subject1Name = "Data structure and Algorithms"
	Subject1Code = "S0001"
	Subject2Name = "Object Oriented Programming"
	Subject2Code = "S0002"
	Subject3Name = "Computer Networks"
	Subject3Code = "S0003"
	Faculty1Code = "F0001"
	Faculty2Code = "F0002"
	Faculty3Code = "F0003"
)

// A stringArray is an array of strings that has been unmarshalled from a JSON
// property that could be either a string or an array of string
type stringArray []string

// SmartContract provides functions for managing a student
type SmartContract struct {
	contractapi.Contract
}

// basic student
type Student struct {
	FirstName          string    `json:"firstName"`
	LastName           string    `json:"lastName"`
	Branch             string    `json:"branch"`
	RegistrationNumber string    `json:"registrationNumber"`
	BloodGroup         string    `json:"bloodGroup"`
	MobileNumber       string    `json:"mobileNumber"`
	Address            string    `json:"address"`
	Subjects           []Subject `json:"subjects"`
	ReportHash         string    `json:"reportHash"`
}

type Subject struct {
	SubjectCode  string  `json:"subjectCode"`
	SubjectName  string  `json:"subjectName"`
	SubjectMarks float64 `json:"subjectMarks"`
	FacultyCode  string  `json:"facultyCode"`
}

type Professor struct {
	FacultyCode string      `json:"facultyCode"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Subjects    stringArray `json:"subject"`
}

// func (sa *stringArray) UnmarshalJSON(data []byte) error {
// 	if len(data) > 0 {
// 		switch data[0] {
// 		case '"':
// 			var s string
// 			if err := json.Unmarshal(data, &s); err != nil {
// 				return err
// 			}
// 			*sa = []string{s}
// 		case '[':
// 			var s []string
// 			if err := json.Unmarshal(data, &s); err != nil {
// 				return err
// 			}
// 			*sa = s
// 		}
// 	}
// 	return nil
// }

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key     string `json:"key"`
	Student *Student
}

// InitLedger adds a base set of student to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	professors := []Professor{
		{
			FacultyCode: Faculty1Code,
			Subjects:    stringArray{Subject1Code},
			FirstName:   "Amit",
			LastName:    "Jaiswal",
		},
		{
			FacultyCode: Faculty2Code,
			Subjects:    stringArray{Subject2Code},
			FirstName:   "Sujit",
			LastName:    "Jaiswal",
		},
		{
			FacultyCode: Faculty3Code,
			Subjects:    stringArray{Subject3Code},
			FirstName:   "Kapil",
			LastName:    "Jaiswal",
		},
	}

	for _, professor := range professors {
		professorAsBytes, _ := json.Marshal(professor)
		if err := ctx.GetStub().PutState(professor.FacultyCode, professorAsBytes); err != nil {
			return fmt.Errorf("Failed to put to world state. %s", professor.FacultyCode)
		}
	}

	subjects1 := []Subject{{FacultyCode: Faculty1Code, SubjectCode: Subject1Code, SubjectName: Subject1Name, SubjectMarks: 80},
		{FacultyCode: Faculty2Code, SubjectCode: Subject2Code, SubjectName: Subject2Name, SubjectMarks: 90},
		{FacultyCode: Faculty3Code, SubjectCode: Subject3Code, SubjectName: Subject3Name, SubjectMarks: 100},
	}

	subjects2 := []Subject{{FacultyCode: Faculty1Code, SubjectCode: Subject1Code, SubjectName: Subject1Name, SubjectMarks: 65},
		{FacultyCode: Faculty2Code, SubjectCode: Subject2Code, SubjectName: Subject2Name, SubjectMarks: 66},
		{FacultyCode: Faculty3Code, SubjectCode: Subject3Code, SubjectName: Subject3Name, SubjectMarks: 67},
	}

	subjects3 := []Subject{{FacultyCode: Faculty1Code, SubjectCode: Subject1Code, SubjectName: Subject1Name, SubjectMarks: 52},
		{FacultyCode: Faculty2Code, SubjectCode: Subject2Code, SubjectName: Subject2Name, SubjectMarks: 55},
		{FacultyCode: Faculty3Code, SubjectCode: Subject3Code, SubjectName: Subject3Name, SubjectMarks: 57},
	}

	students := []Student{
		{FirstName: "Sajan", LastName: "Jaiswal", Branch: "CSE", RegistrationNumber: "1816128", BloodGroup: "A+",
			MobileNumber: "+917064274923", Address: "White House, Motihari, Bihar", Subjects: subjects1,
		},
		{FirstName: "Abhishek", LastName: "Jaiswal", Branch: "CSE", RegistrationNumber: "1816129",
			BloodGroup: "B+", MobileNumber: "+918210791275", Address: "MidLand,Dimapur", Subjects: subjects2,
		},
		{FirstName: "Meso", LastName: "Z", Branch: "CSE", RegistrationNumber: "181630",
			BloodGroup: "B+", MobileNumber: "+918210791275", Address: "MidLand,Dimapur", Subjects: subjects3,
		},
	}

	for _, student := range students {
		studentAsBytes, _ := json.Marshal(student)
		if err := ctx.GetStub().PutState(student.RegistrationNumber, studentAsBytes); err != nil {
			return fmt.Errorf("Failed to put to world state. %s", student.RegistrationNumber)
		}
	}

	return nil
}

// RegisterStudent adds a new Student to the world state with given details
func (s *SmartContract) RegisterStudent(ctx contractapi.TransactionContextInterface, registrationNumber, firstName, lastName, branch, bloodGroup, mobileNumber, address string) error {

	// subjects := stringArray{"Physics", "chemistry", "Math"}
	subjects := []Subject{
		{FacultyCode: Faculty1Code,
			SubjectCode: Subject1Code,
			SubjectName: Subject1Name,
		},
		{FacultyCode: Faculty2Code,
			SubjectCode: Subject2Code,
			SubjectName: Subject2Name,
		},
		{FacultyCode: Faculty3Code,
			SubjectCode: Subject3Code,
			SubjectName: Subject3Name,
		},
	}
	student := Student{
		FirstName:          firstName,
		LastName:           lastName,
		Branch:             branch,
		RegistrationNumber: registrationNumber,
		BloodGroup:         bloodGroup,
		MobileNumber:       mobileNumber,
		Address:            address,
		Subjects:           subjects,
	}

	studentAsBytes, _ := json.Marshal(&student)
	return ctx.GetStub().PutState(student.RegistrationNumber, studentAsBytes)
}

// QueryStudent returns the student stored in the world state with given id
func (s *SmartContract) QueryStudent(ctx contractapi.TransactionContextInterface, registrationNumber string) (*Student, error) {
	cidAsBytes, err := ctx.GetStub().GetState(registrationNumber)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", registrationNumber)
	}
	if cidAsBytes == nil {
		return nil, fmt.Errorf("student %s does not exist", registrationNumber)
	}
	student := new(Student)

	err = json.Unmarshal(cidAsBytes, student)

	if err != nil {
		return nil, fmt.Errorf("failed to parse json %s", err)
	}
	return student, nil
}

// ChangeCarOwner updates the owner field of car with given id in world state
func (s *SmartContract) UpdateMobileNumber(ctx contractapi.TransactionContextInterface, registrationNumber string, mobileNumber string) error {
	student, err := s.QueryStudent(ctx, registrationNumber)

	if err != nil {
		return err
	}
	student.MobileNumber = mobileNumber

	// Converting the map into JSON object
	studentAsBytes, _ := json.Marshal(student)

	return ctx.GetStub().PutState(student.RegistrationNumber, studentAsBytes)

}

func (s *SmartContract) RewardMarks(ctx contractapi.TransactionContextInterface, facultyCode, studentRegistrationNumber, subjectCode string, marks float64) error {

	studentAsBytes, err := ctx.GetStub().GetState(studentRegistrationNumber)
	if err != nil {
		return fmt.Errorf("Failed to read from world state. %s", studentRegistrationNumber)
	}
	if studentAsBytes == nil {
		return fmt.Errorf("student %s does not exist", studentRegistrationNumber)
	}
	student := new(Student)
	if err = json.Unmarshal(studentAsBytes, student); err != nil {
		return fmt.Errorf("failed to parse json %s", err)
	}

	for i, subject := range student.Subjects {
		if subject.FacultyCode == facultyCode && subject.SubjectCode == subjectCode {
			if subject.SubjectMarks == 0 {
				student.Subjects[i].SubjectMarks = marks
			} else {
				return fmt.Errorf("Marks Already rewarded")
			}

		}

	}

	// Converting the map into JSON object
	studentAsBytes, _ = json.Marshal(student)

	return ctx.GetStub().PutState(student.RegistrationNumber, studentAsBytes)
}

func (s *SmartContract) GenrateReport(ctx contractapi.TransactionContextInterface, registrationNumber string) (string, error) {
	studentAsBytes, err := ctx.GetStub().GetState(registrationNumber)
	if err != nil {
		return "", fmt.Errorf("Failed to read from world state. %s", registrationNumber)
	}
	if studentAsBytes == nil {
		return "", fmt.Errorf("student %s does not exist", registrationNumber)
	}
	student := new(Student)

	if err = json.Unmarshal(studentAsBytes, student); err != nil {
		return "", fmt.Errorf("failed to parse json %s", err)
	}

	if student.ReportHash != "" {

		return student.ReportHash, nil
	}

	pdf := newReport(student.FirstName, student.LastName, student.RegistrationNumber)
	if pdf.Err() {
		return "", fmt.Errorf("Failed creating PDF report: %s\n", pdf.Error())
	}

	header1 := []string{"Award Programme", "Bachelor of Computer Science"}

	if pdf = header(pdf, header1); pdf.Err() {
		return "", fmt.Errorf("Failed creating PDF report: %s\n", pdf.Error())
	}
	table1Contents := [][]string{{"Department", "Informatics"}, {"School", "School of Engineering and Informatics"},
		{"Mode of attendance", "Full-time"},
	}

	if pdf = table(pdf, table1Contents); pdf.Err() {
		return "", fmt.Errorf("Failed creating PDF report: %s\n", pdf.Error())
	}
	pdf.Ln(14)
	pdf.SetFont("Times", "B", 18)
	pdf.Cell(80, 10, "Academic Result")
	pdf.Ln(-1)

	header2 := []string{"Module Title", "Marks"}
	if pdf = header(pdf, header2); pdf.Err() {
		return "", fmt.Errorf("Failed creating PDF report: %s\n", pdf.Error())
	}

	totalSum := student.Subjects[0].SubjectMarks + student.Subjects[1].SubjectMarks + student.Subjects[2].SubjectMarks

	outcome := "Fail"

	if totalSum/3 >= 70 {
		outcome = "First-Class Pass"
	}
	if totalSum/3 >= 60 && totalSum/3 <= 69 {
		outcome = "Upper-Second-Class Pass"
	}
	if totalSum/3 >= 50 && totalSum/3 <= 59 {
		outcome = "Lower-Second-Class Pass"
	}
	if totalSum/3 >= 40 && totalSum/3 <= 49 {
		outcome = "Third-Class Pass"
	}
	table2Contents := [][]string{{"Data structure and Algorithms", fmt.Sprintf("%.2f", student.Subjects[0].SubjectMarks)},
		{"Object Oriented Programming", fmt.Sprintf("%.2f", student.Subjects[1].SubjectMarks)},
		{"Computer Networks", fmt.Sprintf("%.2f", student.Subjects[2].SubjectMarks)},
		{"Overall Result", fmt.Sprintf("%.2f", float64(totalSum)/float64(len(student.Subjects))) + "%"},
		{"Outcome", outcome},
	}
	if pdf = table(pdf, table2Contents); pdf.Err() {
		return "", fmt.Errorf("Failed creating PDF report: %s\n", pdf.Error())
	}

	pipeReader, pipeWriter := io.Pipe()

	go func() {
		if err := pdf.OutputAndClose(pipeWriter); err != nil {
			log.Printf("Faile to wirte File to io writer %s", err)
		}
		pipeWriter.Close()

	}()

	reportHash, err := s.uploadReportOnIPFS(pipeReader)
	if err != nil {
		return "", err
	}
	student.ReportHash = reportHash
	studentAsBytes, _ = json.Marshal(student)
	if err := ctx.GetStub().PutState(student.RegistrationNumber, studentAsBytes); err != nil {
		return "", err
	}
	return reportHash, nil
}

func (c *SmartContract) uploadReportOnIPFS(pipeReader *io.PipeReader) (string, error) {

	sh := shell.NewShell(ipfsURL)

	reportHash, err := sh.Add(pipeReader)
	if err != nil {
		return "", fmt.Errorf("Faile to upload file to ipfs %s", err)
	}

	return reportHash, nil
}

// DeleteAsset deletes a given asset from the world state.
func (s *SmartContract) DeleteStudent(ctx contractapi.TransactionContextInterface, registrationNumber string) error {

	student, err := s.QueryStudent(ctx, registrationNumber)
	if err != nil {
		return err
	}

	return ctx.GetStub().DelState(student.RegistrationNumber)
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create student chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting student chaincode: %s", err.Error())
	}

}

func table(pdf *gofpdf.Fpdf, tbl [][]string) *gofpdf.Fpdf {
	// Reset font and fill color.
	pdf.SetFont("Times", "", 16)
	pdf.SetFillColor(255, 255, 255)

	// Every column gets aligned according to its contents.
	align := []string{"L", "C", "L", "R", "R", "R"}
	for _, line := range tbl {
		pdf.CellFormat(100, 7, line[0], "1", 0, align[0], false, 0, "")
		pdf.CellFormat(100, 7, line[1], "1", 0, align[0], false, 0, "")

		pdf.Ln(-1)
	}

	return pdf
}

func newReport(firstName, lastName, registrationNumber string) *gofpdf.Fpdf {

	pdf := gofpdf.New("L", "mm", "Letter", "")

	pdf.AddPage()

	pdf.SetFont("Times", "B", 22)
	pdf.Cell(40, 10, time.Now().Format("Mon Jan 2, 2006"))
	pdf.Ln(14)

	pdf.SetFont("Times", "B", 18)
	pdf.Cell(80, 10, "Student Details")
	pdf.Ln(-1)

	pdf.SetFont("Times", "", 16)
	pdf.Cell(80, 10, "Student ID: "+registrationNumber)
	pdf.Ln(-1)
	pdf.SetFont("Times", "", 16)
	pdf.Cell(80, 10, "First Name: "+firstName)
	pdf.Ln(-1)
	pdf.SetFont("Times", "", 16)
	pdf.Cell(80, 10, "Last Name: "+lastName)
	pdf.Ln(14)

	pdf.SetFont("Times", "B", 18)
	pdf.Cell(80, 10, "Academic Details")
	pdf.Ln(-1)
	return pdf
}

func header(pdf *gofpdf.Fpdf, hdr []string) *gofpdf.Fpdf {
	pdf.SetFont("Times", "B", 16)
	pdf.SetFillColor(240, 240, 240)

	pdf.CellFormat(100, 7, hdr[0], "1", 0, "", true, 0, "")
	pdf.CellFormat(100, 7, hdr[1], "1", 0, "", true, 0, "")

	pdf.Ln(-1)
	return pdf
}
