type Product = {
    id: string,
    name: string,
    price: number
}

type ShoppingCart = {
    products: Product[]
    cartPrice: number
}

type Purchase = {
    person: string
    address: string
    cart: ShoppingCart
}

type CreatPurchaseResponse = {
    id: string
}

type ProductList = Product[]

const api = 'http://localhost:8080/api'

const shoppingCart: ShoppingCart = {products: [], cartPrice: 0}
var productList: ProductList

async function findAllProducts() {
    try {
        const response = await fetch(api + '/products', {
            method: 'GET',
            headers: {
                Accept: 'applicatino/json'
            }
        })

        if (!response.ok) {
            throw new Error('error! status: ${response.status}')
        }

        const result = (await response.json()) as Product[]
        console.log(result)
        productList = result
        return result
    } catch (error) {
        if (error instanceof Error) {
            console.log('error message: ', error.message)
            return error.message
        } else {
            console.log('unexpected error ', error)
            return 'unexpected error'
        }
    }
}

function showProducts(): void {
    const productsElement = document.getElementById('productList')
    const productsPromise = findAllProducts()

    productsPromise.then(products => {
        products = products as ProductList

        products.forEach(product => {
            const productElement = document.createElement('div')
            var priceStr = 'Цена: ' + product.price
            var nameStr = 'Название: ' + product.name

            productElement.textContent = nameStr +'  '+ priceStr +'   '

            productsElement?.appendChild(productElement)

            var toCartButtom = document.createElement('button')
            toCartButtom.textContent = 'в корзину'
            toCartButtom.value = product.id
            toCartButtom.onclick = putProductToShopingCart

            productsElement?.appendChild(toCartButtom)
            productsElement?.appendChild(document.createElement('hr'))
            productsElement?.appendChild(document.createElement('br'))            
        })
    }).catch(err => {
        alert(err)
    })
}

function putProductToShopingCart(event: MouseEvent): void {
    const id = (event.target as HTMLButtonElement).value
    // @ts-ignore
    const product = productList.find((prod) => { return prod.id === id }) as Product
    
    const cartProducts = shoppingCart.products
    // @ts-ignore
    const containsProduct = cartProducts.map(prod => prod.id).includes(product.id)

    if (!containsProduct)  {
        shoppingCart.products.push(product)
        shoppingCart.cartPrice += product.price
        updateShopingCartView()
    }
}

function deleteProductFromShopingCart(event: MouseEvent): void {
    console.log('shoping cart products count before deleting ' + shoppingCart.products.length)
    const id = (event.target as HTMLButtonElement).value
    // @ts-ignore
    const indexOfProd = shoppingCart.products.findIndex((prod) => {
        return prod.id === id
    })
    console.log('index of deleting element ' + indexOfProd)
    if (indexOfProd !== -1) {
        console.log('deleting block start')
        // @ts-ignore
        const deletedProd = shoppingCart.products.find((prod) => { return prod.id === id }) as Product
        console.log(deletedProd.name +'   '+ deletedProd.id)
        shoppingCart.cartPrice -= deletedProd.price
        shoppingCart.products.splice(indexOfProd, 1)
        console.log('delete bloc end. products in shoping cart = ' + shoppingCart.products.length)
        updateShopingCartView()
    }

    console.log('shoping cart products count after deleting ' + shoppingCart.products.length)
}

function updateShopingCartView() {
    console.log('shoping cart products count before showing ' + shoppingCart.products.length)

    const cartElement = document.getElementById('cart-id') as HTMLElement
    console.log('cart-id element count before deleting ' + cartElement.childElementCount)

    while (cartElement.firstChild) {
        cartElement.removeChild(cartElement.firstChild)
    }
    console.log('cart-id element cound after deleting ' + cartElement.childElementCount)

    if (shoppingCart.products.length !== 0) {
        appendToCartElementTable(cartElement)
    }
    
    console.log('shoping cart products count after showing ' + shoppingCart.products.length)
}

function appendToCartElementTable(cartElement: HTMLElement) {
    const cartTable = document.createElement('table')

    const tableHead = document.createElement('tr')

    const name = document.createElement('td')
    name.textContent = 'Продукт'

    const price = document.createElement('td')
    price.textContent = 'Цена'

    const deleteProdButton = document.createElement('td')
    deleteProdButton.textContent = 'действие'

    tableHead.appendChild(name)
    tableHead.appendChild(price)
    tableHead.appendChild(deleteProdButton)

    cartTable.appendChild(tableHead)

    shoppingCart.products.forEach(prod => {
        const prodRow = document.createElement('tr')

        const productName = document.createElement('td')
        productName.textContent = prod.name

        const productPrice = document.createElement('td')
        productPrice.textContent = prod.price.toString()

        const deleteField = document.createElement('td')

        const deleteButton = document.createElement('button')
        deleteButton.value = prod.id
        deleteButton.textContent = 'убрать'
        deleteButton.onclick = deleteProductFromShopingCart

        deleteField.appendChild(deleteButton)

        prodRow.appendChild(productName)
        prodRow.appendChild(productPrice)
        prodRow.appendChild(deleteButton)

        cartTable.appendChild(prodRow)
    })

    cartElement.appendChild(cartTable)

    const cartPrice = document.createElement('div')
    cartPrice.textContent = 'Цена корзины : ' + shoppingCart.cartPrice.toString()

    cartElement.appendChild(cartPrice)
}

async function executePurchase() {
    const personInput = document.getElementById('person-input-id') as HTMLInputElement
    const addressInput = document.getElementById('address-input-id') as HTMLInputElement

    const person = personInput.value
    const address = addressInput.value

    if (!person) {
        alert('Получатель не указан')
        throw new Error("Получатель не указан")
    }

    if (!address) {
        alert('Адрес получателя не указан')
        throw new Error("Адрес получателя не указан")
    }

    if (shoppingCart.products.length === 0) {
        alert('Нужно выбрать как минимум один товар')
        throw new Error('Нужно выбрать как минимум один товар')
    }

    const purchase: Purchase = {
        person: person,
        address: address,
        cart: shoppingCart
    }

    try {
        const response = await fetch(api + '/purchases', {
            method: 'POST',
            body: JSON.stringify(purchase),
            headers: {
                'Content-Type': 'application/json',
                Acept: 'application/json'
            }
        })

        if (!response.ok) {
            alert('Что-то пошло не так')
            throw new Error(`Error! status: ${response.status}`)
        }

        const result = (await response.json()) as CreatPurchaseResponse

        alert('Заказ успешно формлен')
        console.log("Заказ оформлен")
        return result
    } catch(error) {
        if (error instanceof Error) {
            console.log('error message: ', error.message);
            return error.message
        }
        console.log('unexpected error: ', error)
        return 'unexpected error'
    }
}