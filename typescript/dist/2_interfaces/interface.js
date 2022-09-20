"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = (() => {
    function addToCart(item) {
        console.log("[Exercise 2.1]", `Adding "${item.title}" to cart.`);
    }
    addToCart({ id: 1, title: "Concrete shoes" });
    class Person {
        constructor(name, age) {
            this.name = name;
            this.age = age;
        }
    }
    const jane = new Person("Jane", 31);
    console.log("[Exercise 2.2]", `${jane.name} is ${jane.age} years old.`);
    const montreal = {
        coords: {
            latitude: 42.332,
            longitude: -73.324,
        },
        name: "Montréal",
    };
    const tampa = {
        coords: {
            latitude: Number("27.9478"),
            longitude: Number("-82.4584"),
        },
        name: "Tampa",
    };
    function getCityInfo(city) {
        const coords = `(${city.coords.latitude.toFixed(3)}, ${city.coords.longitude.toFixed(3)})`;
        return `${city.name.toUpperCase()} is located at ${coords}.`;
    }
    console.log("[Exercise 2.3]", `${getCityInfo(montreal)} \n\n ${getCityInfo(tampa)}`);
    class User {
        constructor(name, id) {
            this.name = name;
            this.id = id;
        }
    }
    const user = new User("Dog", 1);
    console.log(user.id);
    user.name = "Harold";
    console.log(`User:`, user);
})();
//# sourceMappingURL=interface.js.map