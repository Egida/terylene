package poly

var (
	Title = `

	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠒⠒⠒⠒⠒⠒⠒⠒⢲⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠎⠀⠀⠀⠀⠀⠀⠀⣄⠀⢣⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡜⠀⡸⠁⠀⠀⠀⠀⠀⠘⡄⠀⢣⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡌⠉⠉⠉⠉⠉⠉⠀⡰⠁⠀⢠⣤⡄⣤⠀⠀⠸⡄⠀⠉⠉⠉⠉⠉⠉⢆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡜⠀⢰⠂⠐⠒⠒⠂⠰⡇⠀⠀⠈⠉⠉⠉⠀⠀⠀⣹⠀⠐⠒⠀⠂⠠⡀⠘⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⣀⣀⣀⣀⣀⣀⡼⠀⢠⠃⠀⠀⠀⠀⠀⠀⢹⡀⠀⠀⠀⠀⠀⠀⠀⢠⠃⠀⠀⠀⠀⠀⠀⢷⠀⠘⣄⣀⣀⣀⣀⣀⣀⠀⠀⠀
	⠀⠀⠀⣰⠁⠀⠀⠀⠀⠀⠀⢠⠏⠀⠀⣲⣶⣶⣆⠀⠀⢱⡀⠀⠀⠀⠀⠀⢠⠏⠀⠀⣶⣶⣶⡆⠀⠀⢣⠀⠀⠀⠀⠀⠀⠀⠘⡄⠀⠀
	⠀⠀⣰⠃⢀⠆⠀⠀⠀⠀⠀⢻⡀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠋⠀⠀⠀⠀⠀⢻⡀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠃⠀⠀⠀⠀⠀⢳⠀⠸⡄⠀
	⠀⢠⠃⢀⡜⠀⠀⡀⠀⣀⠀⠀⢣⠀⠀⠀⠀⠀⠀⠀⢀⠏⠀⠀⡀⠀⣀⠀⠀⢣⠀⠀⠀⠀⠀⠀⠀⢠⠃⠀⢀⡀⡀⣀⠀⠀⢃⠀⢱⡀
	⢰⠃⠀⡞⠀⠀⠘⠿⠟⠿⠀⠀⠀⢣⡀⣀⠀⠀⠀⣠⡎⠀⠀⠸⠗⠛⠿⠂⠀⠀⢣⣀⣀⠀⠀⠀⣠⠇⠀⠀⠸⠟⠷⠿⠀⠀⠈⠆⠀⢱
	⠘⣧⠀⢳⠀⠀⠀⠀⠀⠀⠀⠀⢀⡞⠀⠀⠀⠀⠀⠈⢣⠀⠀⠀⠀⠀⠀⠀⠀⢀⡞⠀⠀⠀⠀⠀⠈⣇⠀⠀⠀⠀⠀⠀⠀⠀⢀⠏⠀⡞
	⠀⠉⣆⠈⢣⠀⠀⠀⠀⠀⠀⠀⡞⠀⠀⢤⣤⣤⣄⠀⠈⢆⠀⠀⠀⠀⠀⠀⠀⡞⠀⠀⣤⣤⣤⡄⠀⠙⡄⠀⠀⠀⠀⠀⠀⢀⡞⠀⡼⠁
	⠀⠀⢘⠀⠀⡧⠀⠀⠀⠀⠀⢼⠀⠀⠀⠛⠙⠉⠃⠀⠀⢘⡦⠀⠀⠀⠀⠀⢾⠀⠀⠀⠙⠙⠋⠁⠀⠀⢘⠆⠀⠀⠀⠀⠠⣾⠀⢐⠃⠀
	⠀⢀⡎⠀⣰⠁⠀⠀⠀⠀⠀⠈⣆⠀⠀⠀⠀⠀⠀⠀⠀⡜⠀⠀⠀⠀⠀⠀⠈⣆⠀⠀⠀⠀⠀⠀⠀⠀⡞⠀⠀⠀⠀⠀⠀⠈⡆⠈⢇⠀
	⢀⡞⠀⣰⠃⠀⠀⣦⣶⣶⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀   developed       ⠀⠀⡜⠀⠀⢴⡖⣦⣶⡀⠀⠘⡄⠘⢆
	⢸⡄⠀⢇⠀⠀⠀⠁⠀⠈⠀⠀⠀⡸⠒⠀⠀⠐⠒     By         ⠂⠀⠀⠀⠀⠺⡄⠀⠀⠈⠁⠁⠀⠀⠀⠀⡭⠀⣸
	⠀⠿⡀⠸⣆⠀⠀⠀⠀⠀⠀⠀⣰⠃⠀⠀⠀⠀⠀⠀⢸⣀ Poly       ⠀⠀⠀⠀⠀⠀⠀⠽⡀⠀⠀⠀⠀⠀⠀⠀⣰⠁⣰⠃
	⠀⠀⢳⠀⠘⣄⠀⠀⠀⠀⠀⣠⠃⠀⠀⢽⠿⡿⠿⠀⢿⣿⠿⠷⠷⢷⡿⢿⢿⠷⠷⣷⠂⠀⠀⠀⠀⠀⢱⡀⠀⠀⠀⠀⠀⣰⠃⢠⠇⠀
	⠀⠀⡰⠀⢀⠏⠀⠀⠀⠀⠀⠹⡀⠀⠀⠀⠀⠀⠀⠀⠀⢰⠂⠀⠀⠀⠀⠀⠸⡀⠀⠀⣴⣶⣴⣆⠀⠀⡰⠁⠀⠀⠀⠀⠈⠻⡀⠐⡆⠀
	⠀⣰⠃⢀⠎⠀⠀⣀⣀⣀⠀⠀⠱⡀⠀⠀⠀⠀⠀⠀⢠⠃⠀⢀⡀⣀⢠⠀⠀⢱⡀⠀⠉⠉⠉⠁⠀⢠⠃⠀⢀⣄⣀⣀⠀⠀⢣⠀⠸⡄
	⢰⠃⠀⡞⠀⠀⠀⠛⠛⠛⠀⠀⠀⢳⠀⠀⠀⠀⠀⢠⡇⠀⠀⠘⠛⠛⠛⠂⠀⠀⢳⠀⠀⠀⠀⠀⢰⠃⠀⠀⠙⠋⠛⠛⠀⠀⠀⠧⠀⢱
	⠘⣧⠀⢱⡀⠀⠀⠀⠀⠀⠀⠀⢀⠎⠀⠀⠀⠀⠀⠀⢳⠀⠀⠀⠀⠀⠀⠀⠀⢀⠏⠀⠀⠀⠀⠀⠀⣧⠀⠀⠀⠀⠀⠀⠀⠀⢠⠇⢀⡞
	⠀⠈⣦⠈⢱⠀⠀⠀⠀⠀⠀⢀⡞⠀⠀⢤⣤⣄⣤⠀⠈⢃⠀⠀⠀⠀⠀⠀⢀⡞⠀⠀⣤⣤⣤⡄⠀⠈⢆⠀⠀⠀⠀⠀⠀⢀⠏⠀⡞⠀
	⠀⠀⠘⣆⠀⠓⠀⠀⠀⠀⠀⢺⠀⠀⠀⠉⠈⠉⠁⠀⠀⢈⠇⠀⠀⠀⠀⠀⢾⠀⠀⠀⠉⠉⠉⠁⠀⠀⢸⠆⠀⠀⠀⠀⠀⠋⠀⡜⠁⠀
	⠀⠀⠀⠘⠤⠤⠤⠤⠤⠤⡄⠈⢇⠀⠀⠀⠀⠀⠀⠀⠀⡞⠀⠀⠀⠀⠀⠀⠈⢆⠀⠀⠀⠀⠀⠀⠀⢀⡎⠀⡠⠤⠤⠤⠤⠤⠴⠁⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠹⡀⠈⢆⠀⠀⠀⠀⠀⠀⡜⠀⠀⢰⡿⡷⣿⠀⠀⠈⢆⠀⠀⠀⠀⠀⢀⡎⠀⣰⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠱⡀⠀⠀⠀⠀⠀⠀⠘⢇⠀⠀⠀⠀⠀⠀⠀⠀⠀⡼⠀⠀⠀⠀⠀⠈⠀⢰⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠑⠒⠒⠒⠒⠒⠢⡀⠘⡄⠀⠀⠀⠀⠀⠀⠀⡼⠁⢠⠒⠒⠒⠒⠒⠒⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠱⠀⠸⠄⠀⠀⣀⠀⠀⡴⠁⢀⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢣⠀⠀⠀⠀⠀⠀⠀⠀⢀⡎⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠉⠉⠉⠉⠉⠉⠉⠉⠀⠀⠀

	`
	Welcome = `
	_____               ____ ____  
	|__  /___ _ __ ___  / ___|___ \ 
	  / // _ \ '__/ _ \| |     __) |  beta
	 / /|  __/ | | (_) | |___ / __/ 
	/____\___|_|  \___/ \____|_____|
`

	Cmd = "Z2::->(root)>"

	Help = `
methods  -> list all the methods
help     -> help menu
clear    -> clear screen
list     -> list all bots
transfer -> transfer bots to another zeroC2 server
payload  -> get payload command
exit     -> exit
`
	Methods = `
format:
!method <target> <port> <duration>

             ( う-´)づ︻╦̵̵̿╤── \(˚☐˚”)/
+---------+-----------------------------------------------+
| methods | descriptions                                  |
+---------+-----------------------------------------------+
| UDP     | UDP flood                                     |
| TCP     | TCP flood                                     |
| SYN     | SYN flood                                     |
| HTTP    | HTTP get flood                                |
| UDPRAPE | Crafted UDP flood                             |
| UDP-VIP | UDP flood to bypass discord , amazon and more |
+---------+-----------------------------------------------+

`
	Broadcastmsg = `broadcasted to all terylene`
)
