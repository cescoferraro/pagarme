package types

import (
	"time"

	"github.com/onnidev/api/shared"
	"gopkg.in/mgo.v2/bson"
)

// ClubLead TODO: NEEDS COMMENT INFO
type ClubLead struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate *Timestamp    `bson:"creationDate" json:"creationDate,omitempty"`
	UpdateDate   *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`

	PersonType string `json:"personType" bson:"personType"`
	Done       bool   `json:"done" bson:"done"`
	Stage      string `json:"stage" bson:"stage"`

	AdminName  string `json:"adminName" bson:"adminName"`
	AdminMail  string `json:"adminMail" bson:"adminMail"`
	AdminPhone string `json:"adminPhone" bson:"adminPhone"`

	ClubName        string   `json:"clubName" bson:"clubName"`
	ClubDescription string   `json:"clubDescription" bson:"clubDescription"`
	ClubType        string   `json:"clubType" bson:"clubType"`
	FBURL           string   `json:"fbURL" bson:"fbURL"`
	Image           string   `json:"image" bson:"image"`
	BackgroundImage string   `json:"backgroundImage" bson:"backgroundImage"`
	MusicStyles     []string `json:"musicStyles,omitempty" bson:"musicStyles,omitempty"`

	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	Country string `json:"country" bson:"country"`
	Street  string `json:"street" bson:"street"`
	Number  string `json:"number" bson:"number"`
	Unit    string `json:"unit" bson:"unit"`

	BankCode           string `json:"bankCode" bson:"bankCode"`
	BankBranch         string `json:"bankBranch" bson:"bankBranch"`
	BankBranchVC       string `json:"bankBrasanchVC" bson:"bankBranchVC"`
	BankAccount        string `json:"bankAccount" bson:"bankAccount"`
	BankAccountVC      string `json:"bankAccountVC" bson:"bankAccountVC"`
	BankAccountName    string `json:"bankAccountName" bson:"bankAccountName"`
	BankAccountType    string `json:"bankAccountType" bson:"bankAccountType"`
	BankLegalAddress   string `json:"bankLegalAddress" bson:"bankLegalAddress"`
	BankDocumentNumber string `json:"bankDocumentNumber" bson:"bankDocumentNumber"`
}

// Address sdfkjn
func (lead ClubLead) Address() string {
	return lead.Street + " " + lead.Unit + " " + lead.City + " " + lead.State
}

// UserClub sdfkjn
func (lead ClubLead) UserClub(password string) UserClub {
	horario := Timestamp(time.Now())
	clubid := lead.ID
	mail := shared.NormalizeEmail(lead.AdminMail)
	return UserClub{
		ID:           bson.NewObjectId(),
		CreationDate: &horario,
		Name:         lead.AdminName,
		Mail:         mail,
		Status:       "ACTIVE",
		Password:     shared.EncryptPassword2(password),
		Profile:      "ADMIN",
		Clubs:        []bson.ObjectId{clubid},
	}
}

// RecipientPost sdfkjn
func (lead ClubLead) RecipientPost() RecipientPost {
	return RecipientPost{
		ClubID:     lead.ID.Hex(),
		PersonType: lead.PersonType,
		// Status
		BankCode:        lead.BankCode,
		BankBranch:      lead.BankBranch,
		BankBranchVC:    lead.BankBranchVC,
		BankAccount:     lead.BankAccount,
		BankAccountVC:   lead.BankAccountVC,
		BankAccountName: lead.BankAccountName,
		BankAccountType: lead.BankAccountType,
		DocumentNumber:  lead.BankDocumentNumber,
	}
}

// ClubLeadPostRequest TODO: NEEDS COMMENT INFO
type ClubLeadPostRequest struct {
	AdminName       string `json:"adminName" bson:"adminName"`
	AdminMail       string `json:"adminMail" bson:"adminMail"`
	AdminPhone      string `json:"adminPhone" bson:"adminPhone"`
	Image           string `json:"image" bson:"image"`
	BackgroundImage string `json:"backgroundImage" bson:"backgroundImage"`
}

// ClubLeadPatch skjdnkk
type ClubLeadPatch struct {
	AdminName       string   `json:"adminName,omitempty" bson:"adminName,omitempty"`
	AdminMail       string   `json:"adminMail,omitempty" bson:"adminMail,omitempty"`
	AdminPhone      string   `json:"adminPhone,omitempty" bson:"adminPhone,omitempty"`
	ClubName        string   `json:"clubName,omitempty" bson:"clubName,omitempty"`
	ClubDescription string   `json:"clubDescription,omitempty" bson:"clubDescription,omitempty"`
	ClubType        string   `json:"clubType,omitempty" bson:"clubType,omitempty"`
	FBURL           string   `json:"fbURL,omitempty" bson:"fbURL,omitempty"`
	Image           string   `json:"image,omitempty" bson:"image,omitempty"`
	BackgroundImage string   `json:"backgroundImage,omitempty" bson:"backgroundImage,omitempty"`
	City            string   `json:"city,omitempty" bson:"city,omitempty"`
	State           string   `json:"state,omitempty" bson:"state,omitempty"`
	Country         string   `json:"country,omitempty" bson:"country,omitempty"`
	Street          string   `json:"street,omitempty" bson:"street,omitempty"`
	PersonType      string   `json:"personType,omitempty" bson:"personType,omitempty"`
	Number          string   `json:"number,omitempty" bson:"number,omitempty"`
	Unit            string   `json:"unit,omitempty" bson:"unit,omitempty"`
	Done            bool     `json:"done,omitempty" bson:"done,omitempty"`
	Stage           string   `json:"stage,omitempty" bson:"stage,omitempty"`
	MusicStyles     []string `json:"musicStyles,omitempty" bson:"musicStyles,omitempty"`

	BankCode           string `json:"bankCode,omitempty" bson:"bankCode,omitempty"`
	BankBranch         string `json:"bankBranch,omitempty" bson:"bankBranch,omitempty"`
	BankBranchVC       string `json:"bankBranchVC,omitempty" bson:"bankBranchVC,omitempty"`
	BankAccount        string `json:"bankAccount,omitempty" bson:"bankAccount,omitempty"`
	BankAccountVC      string `json:"bankAccountVC,omitempty" bson:"bankAccountVC,omitempty"`
	BankLegalAddress   string `json:"bankLegalAddress,omitempty" bson:"bankLegalAddress,omitempty"`
	BankAccountName    string `json:"bankAccountName,omitempty" bson:"bankAccountName,omitempty"`
	BankAccountType    string `json:"bankAccountType,omitempty" bson:"bankAccountType,omitempty"`
	BankDocumentNumber string `json:"bankDocumentNumber,omitempty" bson:"bankDocumentNumber,omitempty"`
}
