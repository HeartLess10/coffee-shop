package dbManager

type DbManager interface {
	Connect()
	Create()
	Read()
	Delete()
	Update()
}
