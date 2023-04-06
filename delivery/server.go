package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery/controller"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/delivery/middleware"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
	"log"
)

type Server struct {
	brandUC       usecase.BrandUseCase
	vehicleUC     usecase.VehicleUseCase
	customerUC    usecase.CustomerUseCase
	transactionUC usecase.TransactionUseCase
	engine        *gin.Engine
	host          string
}

func (s *Server) initController() {
	// add middleware
	s.engine.Use(middleware.LogRequestMiddleware())
	controller.NewBrandController(s.engine, s.brandUC)
	controller.NewVehicleController(s.engine, s.vehicleUC)
	controller.NewCustomerController(s.engine, s.customerUC)
	controller.NewTransactionController(s.engine, s.transactionUC)
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		log.Printf("failed to run server :%s", err)
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Printf("failed to serve config :%s", err)
	}

	dbConn, err := config.NewDbConnection(cfg)
	if err != nil {
		log.Printf("failed to serve connection :%s", err)
	}

	db := dbConn.Conn()
	r := gin.Default()
	brandRepo := repository.NewBrandRepository(db)
	vehicleRepo := repository.NewVehicleRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	brandUc := usecase.NewBrandUseCase(brandRepo)
	vehicleUC := usecase.NewVehicleUseCase(vehicleRepo, brandUc)
	customerUC := usecase.NewCustomerUseCase(customerRepo)
	employeeUC := usecase.NewEmployeeUseCase(employeeRepo)
	transactionUC := usecase.NewTransactionUseCase(transactionRepo, vehicleUC, employeeUC, customerUC)
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		brandUC:       brandUc,
		vehicleUC:     vehicleUC,
		customerUC:    customerUC,
		transactionUC: transactionUC,
		engine:        r,
		host:          host,
	}

}
