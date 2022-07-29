type WarRobot struct {
On    bool
Ammo  int
Power int
}

func (wr *WarRobot) Shoot() bool {
if wr.Ammo > 0 && wr.On {
wr.Ammo--
return true
}
return false
}

func (wr *WarRobot) RideBike() bool {
if wr.Power > 0 && wr.On {
wr.Power--
return true
}
return false
}
func main() {

testStruct := new(WarRobot)
/*
 * Экземпляр созданной вами структуры необходимо передать в качестве
 * аргумента функции testStruct, которая выполнит проверку соблюдения
 * всех условий задания/
 */

// }