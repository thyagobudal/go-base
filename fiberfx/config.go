package fiberfx

// Config contém todas as configurações do servidor.
type Config struct {
	Port        string // Porta do servidor
	ServiceName string // Nome do serviço
	Environment string // Ambiente do serviço (ex: development, production)
	EnableAPM   bool   // Habilita ou desabilita o Elastic APM
}
