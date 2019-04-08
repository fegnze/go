-- 自动生成文件，请勿手动修改

local AreaDataTemplateLang = {}

AreaDataTemplateLang["第1章:SIDE3宣战"] = "第1章:SIDE3宣战"
AreaDataTemplateLang["第2章:三号星争夺战"] = "第2章:三号星争夺战"
AreaDataTemplateLang["第3章:月神二号"] = "第3章:月神二号"
AreaDataTemplateLang["第4章:四号星"] = "第4章:四号星"
AreaDataTemplateLang["第5章:目标贾布罗"] = "第5章:目标贾布罗"
AreaDataTemplateLang["第6章:北美大陆"] = "第6章:北美大陆"
AreaDataTemplateLang["第7章:沙漠"] = "第7章:沙漠"
AreaDataTemplateLang["第8章:进军贾布罗"] = "第8章:进军贾布罗"
AreaDataTemplateLang["第9章:抵达贾布罗"] = "第9章:抵达贾布罗"
AreaDataTemplateLang["第10章:雪拉的回忆"] = "第10章:雪拉的回忆"
AreaDataTemplateLang["第11章:过往激战"] = "第11章:过往激战"
AreaDataTemplateLang["第12章:SIDE7"] = "第12章:SIDE7"
AreaDataTemplateLang["第13章:G3毒气"] = "第13章:G3毒气"
AreaDataTemplateLang["第14章:突袭SIDE5"] = "第14章:突袭SIDE5"
AreaDataTemplateLang["第15章:欧洲战场"] = "第15章:欧洲战场"
AreaDataTemplateLang["第16章:重返宇宙"] = "第16章:重返宇宙"
AreaDataTemplateLang["第17章:所罗门战略"] = "第17章:所罗门战略"
AreaDataTemplateLang["第18章:宇宙大决战(上）"] = "第18章:宇宙大决战(上）"
AreaDataTemplateLang["第19章:宇宙大决战(下）"] = "第19章:宇宙大决战(下）"
AreaDataTemplateLang["第20章:战争终结"] = "第20章:战争终结"

function AreaDataTemplateLang.getLang( key, ... )
	return formatLang(AreaDataTemplateLang[key], ...)
end
_G.AreaDataTemplateLang = ReadOnly.readOnly(AreaDataTemplateLang)