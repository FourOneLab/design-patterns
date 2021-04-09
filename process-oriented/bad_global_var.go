package process_oriented

// - 1.把程序中所用的常量都集中存放，影响代码的可维护性，
//     多人开发时，常量数量会更多，修改常量还可能会冲突。
//
// - 2.这样还会增加代码的编译时间,常量越多，依赖这个文件的文件就多，
//     每次修改都会导致依赖它的文件重写编译。
//
// - 3.影响代码的复用性。如果某个项目只依赖其中的一部分常量，
//     也需要把其他不相关的常量一起引入。
//
//  根据常量的用途拆分到不同的包中，这样提高包中类设计的内聚性，和代码复用性。
const (
	MYSQL_ADDR_KEY     = "mysql_addr"
	MYSQL_DB_NAME_KEY  = "db_name"
	MYSQL_USERNAME_KEY = "mysql_username"
	MYSQL_PASSWORD_KEY = "mysql_password"

	REDIS_DEFAULT_ADDR       = "192.168.7.2:7234"
	REDIS_DEFAULT_MAX_TOTAL  = 50
	REDIS_DEFAULT_MAX_IDLE   = 50
	REDIS_DEFAULT_MIN_IDLE   = 20
	REDIS_DEFAULT_KEY_PREFIX = "rt:"
)
