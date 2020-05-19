package main

// This file is automatically generated by github.com/manland/qml-go/cmd/genqrc

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/manland/qml-go"
)

func init() {
	var r *qml.Resources
	var err error
	if os.Getenv("QRC_REPACK") == "1" {
		err = qrcRepackResources()
		if err != nil {
			panic("cannot repack qrc resources: " + err.Error())
		}
		r, err = qml.ParseResources(qrcResourcesRepacked)
	} else {
		r, err = qml.ParseResourcesString(qrcResourcesData)
	}
	if err != nil {
		panic("cannot parse bundled resources data: " + err.Error())
	}
	qml.LoadResources(r)
}

func qrcRepackResources() error {
	subdirs := []string{"assets"}
	var rp qml.ResourcesPacker
	for _, subdir := range subdirs {
		err := filepath.Walk(subdir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			rp.Add(filepath.ToSlash(path), data)
			return nil
		})
		if err != nil {
			return err
		}
	}
	qrcResourcesRepacked = rp.Pack().Bytes()
	return nil
}

var qrcResourcesRepacked []byte
var qrcResourcesData = "qres\x00\x00\x00\x01\x00\x00\t\xeb\x00\x00\x00\x14\x00\x00\t\x9d\x00\x00\x03]\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00 \x00\x00\x00 \b\x06\x00\x00\x00szz\xf4\x00\x00\x00\x19tEXtSoftware\x00Adobe ImageReadyq\xc9e<\x00\x00\x02\xffIDATxڔWM\x8f\xd40\fM\xd2t\x96\x0f\xed\xc2a\xf7\x00\a~\x01\xff\xff\x9fp\xe2\xc0\x05\x84\x84\x90X\x81\xc4\x0e\xdb6&\xae\x9c\xea\xcd\xc3閑\xacd\xda\xd4~v\x9e\x1d'\x87\xff\xf8\x89H\xaaC\x93hc\xd1W&:/1F9\xaa3\x9b\xe2\xbd5jh\xb0\xb5\x99\x9e\a3\x1c`^\xd5\xc9\\\xc7Ť\xaf8\xc6\v\x85=\x80W\xe6i\x00\xcf\x11\x84\x10\x90bϳ\x01\x98\xecY?\x02\x1d\xaf\xd5\xf0\b\x86Q\xdc\x1d\x82mXl\u07b6IA\xccG\x01\xe8G\xcf-\xec\tF\xde{6\x8e\\H\x06\"BĢ\x01\xd9\x05\x80\xc6\a\x02\xd18\x90H1\x86\xbe\x98\xa7\xed}!\x1e\bG\x82\x01\\\x81\xe1\f\xe3\b`\"\xf1\x80\xa3о\xe9\x85]\x10T&0#x<\xd8\xff\x11\xfe\xefeJp\xb6*:\xc63\x80\xdd\x00\xe8\xc2g\x90r͋\x93͙\xf9\t\x14F`y\x04\x0ed\xda\x1a\x01\xe7t\xfe\x88\x002(F\xef\xd9\xf3\x04{ޛ\x17\x00\x96\x81\x98\x02\xeb\xb6\"\xd6\x00\x8c\xf00\x01\xe1\xd8x\xa4T\x8c\x94~b\xdf,\xf0\xae9Q(3V\x00\xc9\xca+\xa7\xda@F\xd8x\xdaIS/:\x19\xb6\xb7\x8dj;採\xc1\xa9zhDS\xf5]\x95[c\xfb\xf7*\x9f\xab\xfc&\x1e\b\x01\xc7\xf3#\x06\nst\x1630\x95\x17U\xdeWym$Ue/\xab\xdcT\xf9P\xe5\x81·\xd4ѹ\xea\xcdNU\x8b\x0e\xeb\x11\xd8[3vm`\x16{\xae9\xff\xa6\xca'\xe2Fp\xf4\xfdC\xac^q\t\x8e\xa2W\x90\xa27\x96\xbe\xad~\\;\aT\xcf\xf8\xb6\x05\xd21\xea\xe5}\x006k\x1e\xff4\xcf'\xc8\xf7@\xa1\x96\x8e\xde\xf5}&/e'\x02M~X\xe8U\xf9\x1f3\xfa`r\x0f \x84\xb6\x82\x8fmi\x00\x8as\xa8`\xe1\xc0:\xaf\n\xbf\x1a\xe9f\x8b`\xb1h\xfc\xaa\xf2\xadcP\xbc\xc6E%\xc3\xc1\x80\v\x17\xe2\a\x82SO?V\xb93>Lf\\S\xf1\xdc1Th\xdc\xdegz\x89G*o\x0f\x16\x97\xb3\xe5\xfd\x17\xa7\x17\x10ꀚ\xbeB\xe7\xc2:\xa6ڗa\aS\x80d3\x9d\xe5\x05\xfa<\x9e\xb3\x0e!\xe33\x91t\xb1\x9ep;\v&\xaa\u05ed\x8dJ\xce\xe9\x17\bX\x8fh\xec\xc8\xe2Da\v\xb3\x17\x81؎L\xa74\xcbN\xca\x16\xe0\xd1#y/d\xebb\x9f\xcfV\xe3g'\rG:\xb2{-\x99P\x9d\xc0\bx\x1c\xb8\xe8\x88\xdaG'\xe8\xeb8\x9c\x83S\xd7y\r\x87\x9d\xe5\xa2/\xe4\x9ep2ţ-j\x8a\x13\x8c^\xe9\xe6L\xf0ȊϺMik\x95\xb0\x9bI\x0e\x00>\xf3\x99|\x9c\x15L\xc2\xdd{A\xe3\xc3\t.&\xec}\x82m\x92N\xd1\x11  \x12\xfd\xd0\xcd(\x00\x89F\xeaf\xc2\x13\xf7C&\x9bx\x17\x92#\x00\x02\x84n\xa0K\x898\x1c\xe8\x95\xde\xf2\xe4\xedXo\xa9G\x80h\x0fG\xc5)\x92\x81\x15@շ\x1c\xbd\x9e\xff\x15`\x00\xdfɄ\xff\xd7\x15\xb9\x95\x00\x00\x00\x00IEND\xaeB`\x82\x00\x00\x06$import QtQuick 2.0\nimport QtQuick.Particles 2.0\nimport QtGraphicalEffects 1.0;\n\nRectangle {\n\tid: root\n\n\twidth: 640\n\theight: 480\n\n\tgradient: Gradient {\n\t\tGradientStop { position: 0.0; color: \"#3a2c32\"; }\n\t\tGradientStop { position: 0.8; color: \"#875864\"; }\n\t\tGradientStop { position: 1.0; color: \"#9b616c\"; }\n\t}\n\n\tText {\n\t\ttext: ctrl.message\n\n\t\tComponent.onCompleted: {\n\t\t\tx = parent.width/2 - width/2\n\t\t\ty = parent.height/2 - height/2\n\t\t}\n\n\t\tcolor: \"white\"\n\t\tfont.bold: true\n\t\tfont.pointSize: 20\n\n\t\tMouseArea {\n\t\t    id: mouseArea\n\t\t    anchors.fill: parent\n\t\t    drag.target: parent\n                    onReleased: ctrl.textReleased(parent)\n\t\t}\n\t}\n\n\tParticleSystem { id: sys }\n\n\tImageParticle {\n\t\tsystem: sys\n\t\tsource: \"qrc:///assets/particle.png\"\n\t\tcolor: \"white\"\n\t\tcolorVariation: 1.0\n\t\talpha: 0.1\n\t}\n\n\tproperty var emitterComponent: Component {\n\t\tid: emitterComponent\n\t\tEmitter {\n\t\t\tid: container\n\t\t\tsystem: sys\n\t\t\tEmitter {\n\t\t\t\tsystem: sys\n\t\t\t\temitRate: 128\n\t\t\t\tlifeSpan: 600\n\t\t\t\tsize: 16\n\t\t\t\tendSize: 8\n\t\t\t\tvelocity: AngleDirection { angleVariation:360; magnitude: 60 }\n\t\t\t}\n\n\t\t\tproperty int life: 2600\n\t\t\tproperty real targetX: 0\n\t\t\tproperty real targetY: 0\n\t\t\temitRate: 128\n\t\t\tlifeSpan: 600\n\t\t\tsize: 24\n\t\t\tendSize: 8\n\t\t\tNumberAnimation on x {\n\t\t\t\tobjectName: \"xAnim\"\n\t\t\t\tid: xAnim;\n\t\t\t\tto: targetX\n\t\t\t\tduration: life\n\t\t\t\trunning: false\n\t\t\t}\n\t\t\tNumberAnimation on y {\n\t\t\t\tobjectName: \"yAnim\"\n\t\t\t\tid: yAnim;\n\t\t\t\tto: targetY\n\t\t\t\tduration: life\n\t\t\t\trunning: false\n\t\t\t}\n\t\t\tTimer {\n\t\t\t\tinterval: life\n\t\t\t\trunning: true\n\t\t\t\tonTriggered: ctrl.done(container)\n\t\t\t}\n\t\t}\n\t}\n}\n\x00\x06\x06\x8a\x9c\xb3\x00a\x00s\x00s\x00e\x00t\x00s\x00\f\b\xf9b\xa7\x00p\x00a\x00r\x00t\x00i\x00c\x00l\x00e\x00.\x00p\x00n\x00g\x00\f\b\xf9e\xdc\x00p\x00a\x00r\x00t\x00i\x00c\x00l\x00e\x00.\x00q\x00m\x00l\x00\x00\x00\x00\x00\x02\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x00\x00\x02\x00\x00\x00\x02\x00\x00\x00\x02\x00\x00\x00\x12\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x000\x00\x00\x00\x00\x00\x01\x00\x00\x03a"
