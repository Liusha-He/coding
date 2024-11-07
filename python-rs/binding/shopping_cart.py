from decimal import Decimal


class ShoppingCart:
    def __init__(self):
        self.items = {}
        self.total = Decimal('0.00')

    def add_item(self, name, price):
        self.items[name] = Decimal(str(price))
        self.total += self.items[name]

    def remove_item(self, name):
        if name in self.items:
            self.total -= self.items[name]
            del self.items[name]

    def get_total(self):
        return self.total

    def get_item_count(self):
        return len(self.items)
