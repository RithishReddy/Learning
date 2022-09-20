"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = (() => {
    function trimmedLength1(text) {
        text;
        if (typeof text === 'string') {
            text;
            return text.trim().length;
        }
        text;
    }
    function trimmedLength2A(text) {
        text;
        if (typeof text === 'string') {
            text;
            return text.trim().length;
        }
        else if (text == null) {
            text;
            return;
        }
        text;
    }
    function trimmedLength2(text) {
        if (text) {
            return text.trim().length;
        }
        text;
    }
    function trimmedLength3(text) {
        if (!text) {
            text;
            return;
        }
        return text.trim().length;
    }
    function trimmedLength4(text) {
        text;
        if (typeof text === 'string') {
            return text.trim().length;
        }
        text;
    }
    console.log('[Exercise 4.0]', `${trimmedLength1("   hi     ")}`);
    function doStuff(value) {
        if (typeof value === 'string') {
            console.log(value.toUpperCase().split('').join(' '));
        }
        else if (typeof value === 'number') {
            console.log(value.toPrecision(5));
        }
        value;
    }
    doStuff(2);
    doStuff(22);
    doStuff(222);
    doStuff('hello');
    console.log('[Exercise 4.1]');
    function padLeft(value, padding) {
        if (padding === 'number') {
            return `${Array(padding + 1).join(' ')}${value}`;
        }
        return padding + value;
    }
    console.log('[Exercise 4.2]', `
    ${padLeft('🐕', 0)}
    ${padLeft('🐕', '🐩')}
    ${padLeft('🐕', '🐩🐩')}
    ${padLeft('🐕', '🐩🐩🐩')}
    ${padLeft('🐕', '🐩🐩🐩🐩')}
  `);
    const numbers = [1, 2, 3, [44, 55], 6, [77, 88], 9, 10];
    function flatten(array) {
        const flattened = [];
        for (const element of array) {
            if (Array.isArray(element)) {
                element;
                flattened.push(...element);
            }
            else {
                element;
                flattened.push(element);
            }
        }
        return flattened;
    }
    const flattenedNumbers = flatten(numbers);
    console.log('[Exercise 4.3]', flattenedNumbers);
    class Bird {
        constructor(species) {
            this.species = species;
        }
        layEggs() {
            console.log('Laying bird eggs.');
        }
        fly(height) {
            console.log(`Flying to a height of ${height} meters.`);
        }
    }
    ;
    class Fish {
        constructor(species) {
            this.species = species;
        }
        layEggs() {
            console.log('Laying fish eggs.');
        }
        swim(depth) {
            console.log(`Swimming to depth of ${depth} meters.`);
        }
    }
    function getRandomAnimal() {
        const animals = [
            new Bird('puffin'),
            new Bird('kittiwake'),
            new Fish('sea robin'),
            new Fish('leafy seadragon'),
        ];
        return animals[Math.floor(Math.random() * animals.length)];
    }
    function interrogateAnimal(animal = getRandomAnimal()) {
        if (animal instanceof Fish)
            animal.swim(10);
        else
            animal.fly(10);
        return animal.species;
    }
    console.log('[Exercise 4.4]', `We've got a ${interrogateAnimal()} on our hands!`);
})();
//# sourceMappingURL=types.js.map