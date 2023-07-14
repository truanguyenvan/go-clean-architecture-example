package config

var DefaultConfig = Configuration{
	Authorization: Authorization{
		CasbinModelFilePath:  "./config/casbin/rbac_model.config",
		CasbinPolicyFilePath: "./config/casbin/policy.csv",
	},
}
