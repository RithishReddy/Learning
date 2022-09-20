"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = (() => {
    class MC {
        greet(event = 'party') {
            return `Welcome to the ${event}`;
        }
    }
    const mc = new MC();
    console.log('[Exercise 5.1]', mc.greet('show'));
    class Person {
        constructor(name, age) {
            this.name = name;
            this.age = age;
        }
    }
    const jane = new Person('Jane', 31);
    console.log('[Exercise 5.2]', `The new person's name is ${jane.name}.`);
    class Employee {
        constructor(title, salary) {
            this.title = title;
            this.salary = salary;
        }
    }
    const employee = new Employee('Engineer', 100000);
    console.log('[Exercise 5.3]', `The new employee's title is ${employee.title} and they earn $ ${employee.salary}.`);
    class Animal {
        constructor(name) {
            this.name = name;
            console.log("In animal", name);
        }
        move(meters) {
            console.log(`${this.name} moved ${meters}m.`);
        }
    }
    class Snake extends Animal {
        constructor(name) {
            super(name);
            this.name = name;
        }
        move(meters = 5) {
            console.log('Slithering...');
            super.move(meters);
        }
    }
    class Pony extends Animal {
        constructor(name) {
            super(name);
            this.name = name;
        }
        move(meters) {
            console.log('Galloping...');
            super.move(meters);
        }
    }
    const sammy = new Snake("Sammy the Snake");
    sammy.move();
    const pokey = new Pony("Pokey the Pony");
    pokey.move(34);
    class Furniture {
        constructor(manufacturer = 'IKEA') {
            this.manufacturer = manufacturer;
        }
    }
    class Desk extends Furniture {
        kind() {
            console.log('[Exercise 5.5]', `This is a desk made by ${this.manufacturer}`);
        }
    }
    class Chair extends Furniture {
        kind() {
            console.log('[Exercise 5.5]', `This is a chair made by ${this.manufacturer}`);
        }
    }
    const desk = new Desk();
    desk.kind();
    desk.manufacturer;
    const chair = new Chair();
    chair.kind();
    chair.manufacturer;
    class Student {
        constructor(name) {
            this.name = name;
        }
        ;
        introduction() {
            console.log('[Exercise 5.6]', `Hi, my name is ${this.name} and I attend ${Student.school}`);
        }
    }
    Student.school = 'Harry Herpson High School';
    const student = new Student('Morty');
    console.log(Student.school);
    student.introduction();
})();
//# sourceMappingURL=classes.js.map