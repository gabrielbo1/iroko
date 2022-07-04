package pkg

type Module struct {
	ModuleName string `json:"module_name"`
	Routes     []Route
}

var modules []Module

// Add module to start application.
func AddModule(module *Module) {
	if modules == nil {
		modules = []Module{}
	}
	for _, m := range modules {
		if m.ModuleName == module.ModuleName {
			return
		}
	}
	modules = append(modules, *module)
}

// Get all modules.
func GetModules() []Module {
	return modules
}
