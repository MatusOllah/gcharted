package funkin

type Section struct {
	SectionNotes   [][]float64 `json:"sectionNotes"`
	LengthInSteps  int         `json:"lengthInSteps"`
	TypeOfSection  int         `json:"typeOfSection"`
	MustHitSection bool        `json:"mustHitSection"`
	Bpm            int         `json:"bpm"`
	ChangeBPM      bool        `json:"changeBPM"`
	AltAnim        bool        `json:"altAnim,omitempty"`
}
