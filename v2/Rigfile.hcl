#
# log.Printf("%s", viper.AllSettings())
# tool := viper.Get("tool")
#
# log.Printf("%s", tool)
#
# viper.Unmarshal(&rig)
#
# log.Printf("[INFO] %s", rig)
# log.Printf("[INFO] %s", rig.Tool[0])
# log.Printf("[INFO] %s", rig.Tool[0]["helm"][0].Url)
#
# type Rigfile struct {
#   Project string
#   Tool []map[string][]Tool
# }
#
# type Tool struct {
#   Url string
# }


# project = "seth"
#
# tool "helm" {
#   url = "test"
# }
#
# tool "test2" {
#   url = "test2"
# }
