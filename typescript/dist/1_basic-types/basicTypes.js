"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = (() => {
    let pi = '3.14159';
    let tau = pi * 2;
    console.log(typeof (tau));
    console.log('[Exercise 1.1]', `${tau} is ${pi} times two.`);
    let pie;
    pie = 'blueberry';
    console.log('[Exercise 1.2]', `I like to eat ${pie}-flavored pie.`);
    let isMark = true;
    console.log('[Exercise 1.3]', `${isMark ? 'Oh, hi Mark' : 'Who are you?'}`);
    const integer = 6;
    const float = 6.66;
    const hex = 0xf00d;
    const binary = 0b1010011010;
    const octal = 0o744;
    const negZero = -0;
    const actuallyNumber = NaN;
    const largestNumber = Number.MAX_VALUE;
    const mostBiglyNumber = Infinity;
    const members = [
        integer,
        float,
        hex,
        binary,
        octal,
        negZero,
        actuallyNumber,
        largestNumber,
        mostBiglyNumber
    ];
    members[0] = '12345';
    console.log('[Exercise 1.4]', members);
    const sequence = Array.from(Array(10).keys());
    const animals = ['pangolin', 'aardvark', 'echidna', 'binturong'];
    const stringsAndNumbers = [1, 'one', 2, 'two', 3, 'three'];
    const allMyArrays = [sequence, animals, stringsAndNumbers];
    console.log('Exercise 1.5', allMyArrays);
    const inventoryItem = ['fidget wibbit', 11];
    const [name, qty] = inventoryItem;
    const msg = addInventory(name, qty);
    console.log('[Exercise 1.6]', msg);
    function addInventory(name, quantity) {
        return `Added ${quantity} ${name}s to inventory.`;
    }
})();
//# sourceMappingURL=basicTypes.js.map