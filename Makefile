
.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module  ${ROOT_MOD}/app/frontend --idl ../../idl/frontend/auth_page.proto
	cwgo server --type HTTP --idl  ../../idl/frontend/auth_page.proto  --service frontend -module gomall/app/frontend -I ../../idl/


.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --idl  ../idl/cart.proto  --service cart --I ../idl/ --module gomall.local/rpc_gen
	@cd app/cart && cwgo server --type RPC --idl  ../../idl/cart.proto --service cart --I ../../idl/ --module gomall/app/cart --pass "-use gomall.local/rpc_gen/kitex_gen"

.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --idl  ../idl/product.proto  --service product --I ../idl/ --module gomall.local/rpc_gen
	@cd app/product && cwgo server --type RPC --idl  ../../idl/product.proto --service product --I ../../idl/ --module gomall/app/product --pass "-use gomall.local/rpc_gen/kitex_gen"

.PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --idl  ../idl/payment.proto  --service payment --I ../idl/ --module gomall.local/rpc_gen
	@mkdir -p app/payment && cd app/payment && cwgo server --type RPC --idl  ../../idl/payment.proto --service payment --I ../../idl/ --module gomall/app/payment --pass "-use gomall.local/rpc_gen/kitex_gen"

.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --idl  ../idl/checkout.proto  --service checkout --I ../idl/ --module gomall.local/rpc_gen
	@mkdir -p app/checkout && cd app/checkout && cwgo server --type RPC --idl  ../../idl/checkout.proto --service checkout --I ../../idl/ --module gomall/app/checkout --pass "-use gomall.local/rpc_gen/kitex_gen"cwgo server --type HTTP --idl  ../../idl/frontend/auth_page.proto  --service frontend -module gomall/app/frontend -I ../../idl/