//PlatForm abstract stb hw modules and give interfaces for logic module use
//Basic hw init and os(kernel) should NOT be part of this module.
//

package platform

import (
	"fmt"
	//"os"
	"errors"
	//"encoding/json"
)

//InitJSONErr shows init json file error
var InitJSONErr = errors.New("Error in init json file")


//HwModuler is the main module in stb
type HwModuler interface {
	//HwModInit should be purecoded and multi-called
	//p *os.File should be stb describe json file
	HwModInit(hw *HwModule) (err error)
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
//PlatForm is the Set of hwmodules and unique id
//to abstract the plaftorm
type PlatForm struct {
	stbid string 
	stbversion int
	modules []HwModule
}

//HwModInit returns 0 means success parse the file and init submodules.
//HwModInit returns line num >0 and errors in which module is error.
func HwModInit(hw *HwModule) (err error) {
	fmt.Println(hw)
	return err
}

//PlatInit returns 0 means success parse the file and init submodules.
//PlatInit returns line num >0 and errors in which module is error.
func PlatInit(plat *PlatForm) (err error) {
	fmt.Println(plat)
	return err
}
