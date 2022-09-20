"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = (() => {
    function add(x, y) {
        return x + y;
    }
    function sumArray(numbers) {
        return numbers.reduce(add, 0);
    }
    const someSum = sumArray([3, 6, 9]);
    console.log('[Exercise 3.1]', `3 + 6 + 9 === ${someSum}`);
    const bankAccount = {
        money: 0,
        deposit(value, message) {
            this.money += value;
            if (message) {
                console.log(message);
            }
        }
    };
    bankAccount.deposit(20);
    bankAccount.deposit(10, 'Deposit received');
    console.log('[Exercise 3.2]', `Account value: $${bankAccount.money}`);
    function computeScore(word) {
        const letters = word.toUpperCase().split('');
        return letters.reduce((accum, curr) => accum += getPointsFor(curr), 0);
    }
    function getPointsFor(letter) {
        const lettersAndPoints = [
            ['AEOIULNRST', 1],
            ['DG', 2],
            ['BCMP', 3],
            ['FHVWY', 4],
            ['K', 5],
            ['JX', 8],
            ['QZ', 10],
        ];
        return lettersAndPoints.reduce((computedScore, pointsTuple) => {
            const [letters, score] = pointsTuple;
            if (letters.split('').find((ll) => ll === letter)) {
                return computedScore += score;
            }
            return computedScore;
        }, 0);
    }
    console.log('[Exercise 3.3]', `zoo is worth ${computeScore('zoo')} points.`);
    function greet(greeting = "hello") {
        return greeting.toUpperCase();
    }
    const defaultGreeting = greet();
    const portugueseGreeting = greet('Oi como vai!');
    console.log('[Exercise 3.4]', defaultGreeting, portugueseGreeting);
    function layEggs(quantity, color) {
        console.log(`[Exercise 3.5] You just laid ${quantity} ${color} eggs. Good job!`);
    }
    layEggs(4, "white");
    let multiply;
    let capitalize;
    capitalize = function (value) {
        return `${value.charAt(0).toUpperCase()}${value.slice(1)}`;
    };
    multiply = function (x, y) {
        return x * y;
    };
    console.log('[Exercise 3.6]', capitalize(`nifty ${multiply(5, 10)}`));
    const numberCollection = [];
    const stringCollection = [];
    function pushToCollection(item, collection) {
        collection.push(item);
        return collection;
    }
    pushToCollection('hi', stringCollection);
    pushToCollection(1, numberCollection);
    pushToCollection(2, numberCollection);
    pushToCollection(3, numberCollection);
    const incrementedByTwo = numberCollection.map((num) => num + 2);
    console.log('[Exercise 3.7]', `[${incrementedByTwo}] should deeply equal [3,4,5]`);
})();
//# sourceMappingURL=functions.js.map