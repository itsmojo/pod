package response

import (
	"encoding/hex"
)

type Type50StatusResponse struct {
	Seq uint16
}

func (r *Type50StatusResponse) Marshal() ([]byte, error) {
	response, _ := hex.DecodeString("02CB500F184C902A005190320054912A80599033805C902B80019133BF04912ABF099033BF0C312ABF119033BF149029BF199134BF1C91283F2190323F24902A3F2990333F2C912A3F3190353F34912A3F3991353F3C312A3F4191363F44902A3F4931343F4C31293F5131343F5431293F5931353F5C312ABF613035BF643029BF0131338104312981093033810C302A811130308114312A81193031011C302A012131320124312A01293133012C312A013130320134312A01393131013C312B014130320144312B0149303001")

	return response, nil
}