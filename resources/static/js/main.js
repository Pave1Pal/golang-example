var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
var api = 'http://localhost:8080/api';
var shoppingCart = { products: [], cartPrice: 0 };
var productList;
function findAllProducts() {
    return __awaiter(this, void 0, void 0, function () {
        var response, result, error_1;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    _a.trys.push([0, 3, , 4]);
                    return [4 /*yield*/, fetch(api + '/products', {
                            method: 'GET',
                            headers: {
                                Accept: 'applicatino/json'
                            }
                        })];
                case 1:
                    response = _a.sent();
                    if (!response.ok) {
                        throw new Error('error! status: ${response.status}');
                    }
                    return [4 /*yield*/, response.json()];
                case 2:
                    result = (_a.sent());
                    console.log(result);
                    productList = result;
                    return [2 /*return*/, result];
                case 3:
                    error_1 = _a.sent();
                    if (error_1 instanceof Error) {
                        console.log('error message: ', error_1.message);
                        return [2 /*return*/, error_1.message];
                    }
                    else {
                        console.log('unexpected error ', error_1);
                        return [2 /*return*/, 'unexpected error'];
                    }
                    return [3 /*break*/, 4];
                case 4: return [2 /*return*/];
            }
        });
    });
}
function showProducts() {
    var productsElement = document.getElementById('productList');
    var productsPromise = findAllProducts();
    productsPromise.then(function (products) {
        products = products;
        products.forEach(function (product) {
            var productElement = document.createElement('div');
            var priceStr = '????????: ' + product.price;
            var nameStr = '????????????????: ' + product.name;
            productElement.textContent = nameStr + '  ' + priceStr + '   ';
            productsElement === null || productsElement === void 0 ? void 0 : productsElement.appendChild(productElement);
            var toCartButtom = document.createElement('button');
            toCartButtom.textContent = '?? ??????????????';
            toCartButtom.value = product.id;
            toCartButtom.onclick = putProductToShopingCart;
            productsElement === null || productsElement === void 0 ? void 0 : productsElement.appendChild(toCartButtom);
            productsElement === null || productsElement === void 0 ? void 0 : productsElement.appendChild(document.createElement('hr'));
            productsElement === null || productsElement === void 0 ? void 0 : productsElement.appendChild(document.createElement('br'));
        });
    })["catch"](function (err) {
        alert(err);
    });
}
function putProductToShopingCart(event) {
    var id = event.target.value;
    // @ts-ignore
    var product = productList.find(function (prod) { return prod.id === id; });
    var cartProducts = shoppingCart.products;
    // @ts-ignore
    var containsProduct = cartProducts.map(function (prod) { return prod.id; }).includes(product.id);
    if (!containsProduct) {
        shoppingCart.products.push(product);
        shoppingCart.cartPrice += product.price;
        updateShopingCartView();
    }
}
function deleteProductFromShopingCart(event) {
    console.log('shoping cart products count before deleting ' + shoppingCart.products.length);
    var id = event.target.value;
    // @ts-ignore
    var indexOfProd = shoppingCart.products.findIndex(function (prod) {
        return prod.id === id;
    });
    console.log('index of deleting element ' + indexOfProd);
    if (indexOfProd !== -1) {
        console.log('deleting block start');
        // @ts-ignore
        var deletedProd = shoppingCart.products.find(function (prod) { return prod.id === id; });
        console.log(deletedProd.name + '   ' + deletedProd.id);
        shoppingCart.cartPrice -= deletedProd.price;
        shoppingCart.products.splice(indexOfProd, 1);
        console.log('delete bloc end. products in shoping cart = ' + shoppingCart.products.length);
        updateShopingCartView();
    }
    console.log('shoping cart products count after deleting ' + shoppingCart.products.length);
}
function updateShopingCartView() {
    console.log('shoping cart products count before showing ' + shoppingCart.products.length);
    var cartElement = document.getElementById('cart-id');
    console.log('cart-id element count before deleting ' + cartElement.childElementCount);
    while (cartElement.firstChild) {
        cartElement.removeChild(cartElement.firstChild);
    }
    console.log('cart-id element cound after deleting ' + cartElement.childElementCount);
    if (shoppingCart.products.length !== 0) {
        appendToCartElementTable(cartElement);
    }
    console.log('shoping cart products count after showing ' + shoppingCart.products.length);
}
function appendToCartElementTable(cartElement) {
    var cartTable = document.createElement('table');
    var tableHead = document.createElement('tr');
    var name = document.createElement('td');
    name.textContent = '??????????????';
    var price = document.createElement('td');
    price.textContent = '????????';
    var deleteProdButton = document.createElement('td');
    deleteProdButton.textContent = '????????????????';
    tableHead.appendChild(name);
    tableHead.appendChild(price);
    tableHead.appendChild(deleteProdButton);
    cartTable.appendChild(tableHead);
    shoppingCart.products.forEach(function (prod) {
        var prodRow = document.createElement('tr');
        var productName = document.createElement('td');
        productName.textContent = prod.name;
        var productPrice = document.createElement('td');
        productPrice.textContent = prod.price.toString();
        var deleteField = document.createElement('td');
        var deleteButton = document.createElement('button');
        deleteButton.value = prod.id;
        deleteButton.textContent = '????????????';
        deleteButton.onclick = deleteProductFromShopingCart;
        deleteField.appendChild(deleteButton);
        prodRow.appendChild(productName);
        prodRow.appendChild(productPrice);
        prodRow.appendChild(deleteButton);
        cartTable.appendChild(prodRow);
    });
    cartElement.appendChild(cartTable);
    var cartPrice = document.createElement('div');
    cartPrice.textContent = '???????? ?????????????? : ' + shoppingCart.cartPrice.toString();
    cartElement.appendChild(cartPrice);
}
function executePurchase() {
    return __awaiter(this, void 0, void 0, function () {
        var personInput, addressInput, person, address, purchase, response, result, error_2;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    personInput = document.getElementById('person-input-id');
                    addressInput = document.getElementById('address-input-id');
                    person = personInput.value;
                    address = addressInput.value;
                    if (!person) {
                        alert('???????????????????? ???? ????????????');
                        throw new Error("???????????????????? ???? ????????????");
                    }
                    if (!address) {
                        alert('?????????? ???????????????????? ???? ????????????');
                        throw new Error("?????????? ???????????????????? ???? ????????????");
                    }
                    if (shoppingCart.products.length === 0) {
                        alert('?????????? ?????????????? ?????? ?????????????? ???????? ??????????');
                        throw new Error('?????????? ?????????????? ?????? ?????????????? ???????? ??????????');
                    }
                    purchase = {
                        person: person,
                        address: address,
                        cart: shoppingCart
                    };
                    _a.label = 1;
                case 1:
                    _a.trys.push([1, 4, , 5]);
                    return [4 /*yield*/, fetch(api + '/purchases', {
                            method: 'POST',
                            body: JSON.stringify(purchase),
                            headers: {
                                'Content-Type': 'application/json',
                                Acept: 'application/json'
                            }
                        })];
                case 2:
                    response = _a.sent();
                    if (!response.ok) {
                        alert('??????-???? ?????????? ???? ??????');
                        throw new Error("Error! status: ".concat(response.status));
                    }
                    return [4 /*yield*/, response.json()];
                case 3:
                    result = (_a.sent());
                    alert('?????????? ?????????????? ??????????????');
                    console.log("?????????? ????????????????");
                    return [2 /*return*/, result];
                case 4:
                    error_2 = _a.sent();
                    if (error_2 instanceof Error) {
                        console.log('error message: ', error_2.message);
                        return [2 /*return*/, error_2.message];
                    }
                    console.log('unexpected error: ', error_2);
                    return [2 /*return*/, 'unexpected error'];
                case 5: return [2 /*return*/];
            }
        });
    });
}
