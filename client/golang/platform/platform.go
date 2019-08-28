//PlatForm abstract stb hw modules and give interfaces for logic module use
//Basic hw init and os(kernel) should NOT be part of this module.
//

package platform

import (
	"os"
	"errors"
	"encoding/json"
)

//InitJSONErr shows init json file error
var InitJSONErr = errors.New("Error in init json file")


//HwModuler is the main module in stb
type HwModuler interface {
	//HwModInit should be purecoded and multi-called
	//p *os.File should be stb describe json file
	HwModInit(p *os.File) (n int, err error)
}

//HwModule is one hwmodule
type HwModule struct {
	name string
	hwtype string
	//is dynamic meas hwmodule will insert/eject in runtime(true)
	isdynamic bool
	//is singleinst means only one instance use hw module(true) one time
	issingleinst bool
	//absolute path in kernel
	hwpath string
	//hw specify attr string list
	hwattr string
}
//PlatForm give the  
//
type PlatForm struct {
	stbid string 
	stbversion int
	modules []HwModule
}

//HwModInit returns 0 means success parse the file and init submodules.
//HwModInit returns line num >0 and errors in which module is error.
func (hwmod *HwModule)HwModInit(p *os.File) (n int, err error) {
	var modules []HwModule
	finfo, _:= p.Stat()
	data := make([]byte, finfo.Size())
	err = json.Unmarshal(data, &modules)
	return 0, err
}
