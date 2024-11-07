import pytest
from pytest_bdd import scenario, given, when, then, parsers

from decimal import Decimal

from shopping_cart import ShoppingCart


@pytest.fixture
def cart():
    return ShoppingCart()


# scenario I
@scenario('features/cart.feature', 'Add item to empty cart')
def test_add_item():
    pass


@given('the cart is empty')
def empty_cart():
    return ShoppingCart()


@when(parsers.parse('I add a "{item}" with price {price:f}'))
def add_item_to_cart(cart, item, price):
    cart.add_item(item, price)


@then(parsers.parse('the cart should contain {count:d} item'))
def check_cart_size(cart, count):
    assert cart.get_item_count() == count

# scenario II
@scenario('features/cart.feature', 'Remove item from cart')
def test_remove_item():
    pass


@given(parsers.parse('I have a "{item}" in my cart with price {price:f}'))
def cart_with_item(item, price):
    cart = ShoppingCart()
    cart.add_item(item, price)
    return cart


@when(parsers.parse('I remove the "{item}" from the cart'))
def remove_item_from_cart(cart, item):
    cart.remove_item(item)


@then('the cart should be empty')
def check_cart_empty(cart):
    assert cart.get_item_count() == 0


@then(parsers.parse('the total should be {total:f}'))
def check_cart_total(cart, total):
    assert cart.get_total() == Decimal(str(total))
