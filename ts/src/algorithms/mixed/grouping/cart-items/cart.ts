export let cartItems = [
  {id: 1, name: "Mango", price: 20},
  {id: 2, name: "Banana", price: 30},
  {id: 3, name: "Apple", price: 40},
  {id: 3, name: "Apple", price: 40},
  {id: 3, name: "Apple", price: 40},
  {id: 3, name: "Apple", price: 40},
];

type Fruit = (typeof cartItems)[number] & {qty?: number};
type GroupedFruit = {
  id: number;
  price: number;
  qty: number;
};

export function groupFruits(fruits: Fruit[]): Record<string, GroupedFruit> {
  return fruits.reduce<Record<string, GroupedFruit>>((obj, item) => {
    if (!obj[item.name]) {
      obj[item.name] = {id: item.id, price: item.price, qty: 1};
    } else {
      let x = obj[item.name];
      obj[item.name] = {...x, qty: (x.qty || 0) + 1};
    }
    return obj;
  }, {});
}

export function groupFruitsV2(fruits: Fruit[]): Record<string, GroupedFruit> {
  return fruits.reduce<Record<string, GroupedFruit>>((obj, fruit) => {
    const {id, price, name} = fruit;
    const existingFruit = obj[name];

    obj[name] = existingFruit
      ? {...existingFruit, qty: existingFruit.qty + 1}
      : {id, price, qty: 1};

    return obj;
  }, {});
}

export function groupToList(items: Fruit[]) {
  return items.reduce<
    {
      id: number;
      name: string;
      price: number;
      qty: number;
    }[]
  >((xs, item) => {
    let foundItem = xs.find((x) => x.id === item.id);
    if (foundItem) {
      foundItem.qty = foundItem.qty += 1;
    } else {
      xs.push({...item, qty: 1});
    }
    return xs;
  }, []);
}
