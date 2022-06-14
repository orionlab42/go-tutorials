Ch1 - ES06
-----------------------

# 1.1 Let vs Var vs Const
* Variable declared with "var" are scoped to the function.
* Variables declared with "let" or "const" are scoped to the block in which they are defined.

# 1.2 Objects
* Objects are key, value pairs.

```jsx
const person = {
    name: 'Mosh',
    walk: function() {},
    talk() {}
};

person.talk();
person['name'] = "John"; // this notation is usefull, when we don't know the name of the key, like in the following ex.:
const targetMember = 'name';
person[targetMember.value] = "Helen";
```

# 1.3 The "this" keyword
* "this" returns a reference to the current object if it is called as a method.
* If you call the function as a standalone function or outside an object "this" returns a reference to global object, 
* which is the window object.

```jsx
const person = {
    name: 'Mosh',
    talk() {
        console.log(this);
    }
};

person.walk(); // returns the person object

const walk = person.walk;
walk();  // returns undefined instead of the window object because strict mode of react is enabled.

```

# 1.4 Binding "this" 
* In JS functions are objects.
* With the bind method we can set the value of this permanently.

```jsx
const person = {
name: 'Mosh',
talk() {
console.log(this);
}
};

person.walk(); // returns the person object

const walk = person.walk.bind(person);
walk();  // returns the person object
```


# 1.5 Arrow functions

```jsx
const square1 = function(number) {
    return number * number;
}

const square2 = () => { // in case there are no parameter
    return number * number;
}

const square3 = number => {
    return number * number;
}

const square4 = number => number * number; // if we return only one line of code, this is the short way

const jobs = [
    {id:1, isActive: true},
    {id:2, isActive: true},
    {id:3, isActive: false},
];
const activeJobs1 = jobs.filer(function(job) { return job.isActive; });
const activeJobs2 = jobs.filer(job =>  job.isActive);
```

# 1.6 Arrow functions and "this"
* setTimeout executes a function after a certain time, in the ex. is 1 min.
* Arrow functions don't rebind the "this" keyword.

```jsx
const person = {
    talk() {
        setTimeout(function() {
            console.log("This", this);
        }, 1000);
    }
};
person.talk(); // returns a reference to the window object, because here inside the setTimeout we have a 
// standalone function.

const person2 = {
    talk() {
        var self = this;
        setTimeout(function() {
            console.log("Self", self);
        }, 1000);
    }
};
person2.talk(); // returns a reference to the person2 object

const person3 = {
    talk() {
        setTimeout(() => {
            console.log("This", this);
        }, 1000);
    }
};
person2.talk(); // returns a reference to the person3 object
```

# 1.7 Array.map method
* It is used to render an array, takes a callback function to do something with each item of the array.

```jsx
const colors = ['red', 'green', 'blue'];
const items = colors.map(function(color) { 
    return '<li>' + color + '</li>';
});
const items2 = colors.map(color => `<li>${color}</li>`); // this is what we call a template literals
```

# 1.8 Object Destructuring

```jsx
const address = {
    street: '',
    city: '',
    country: ''
};
const street1 = address.street;
const city1 = address.city;
const country1 = address.country;

const {street, city, country} = address;
const {street} = address; // you can also take only one property
const {street: st} = address; // using an alias
```

# 1.9 Spread Operator
* Works with arrays and objects.

```jsx
const first = [1, 2, 3];
const second = [4, 5, 6];
const combined = first.concat(second); // the old way
const combined2 = [...first, 'a', ...second, 'b'];
const clone = [...first]; // we created a clone of first

const first1 = {name: 'Mosh'};
const second1 = { job: 'Instructor'};
const combined1 = {...first1, ...second1, location: 'Australia'}
```

# 1.10 Classes
* Classes are a blueprint for objects.

```jsx
class Person {
    constructor(name) {
        this.name = name; // this initializes the name property
    }
    
    walk() {
        console.log('walk');
    }
}

const person = new Person('Mosh');
```

# 1.11 Inheritance
* With the "extend" of a class inherits all the method of another class.

```jsx
class Person {
    constructor(name) {
        this.name = name; // this initializes the name property
    }
    
    walk() {
        console.log('walk');
    }
}

class Teacher extends Person{
    constructor(name, degree) {
        super(name); // references the parent class
        this.degree = degree;
    }
    
    teach() {
        console.log("teach");
    }
}
const teacher = new Teacher('Mosh', 'Msc');
teacher.walk();
```

# 1.12 Modules, Named and Default Exports
* We can make a function public or visible outside a file by exporting (prefix the class with export)
* What is exported has a name, and it can be imported.
* Default --> import ... from '';
* Named --> import {...} from '';

```jsx
// file1
export function promote() {}

export class Person {
    constructor(name) {
        this.name = name; // this initializes the name property
    }
    
    walk() {
        console.log('walk');
    }
}

// file2
import { Person, promote }  from './file1';
export function promote2() {}
export default class Teacher extends Person{
    constructor(name, degree) {
        super(name); // references the parent class
        this.degree = degree;
    }
    
    teach() {
        console.log("teach");
    }
}

// index file
import Teacher, { promote2 }  from './file2';
const teacher = new Teacher('Mosh', 'Msc');
teacher.walk();
```

Ch2 - Components
-----------------------

# 2.1 Basics

```jsx
class Counter extends Component {
    state = {
        count: 1,
        tags: ["tag1", "tag2", "tag3"]
    };
    
    //adding style
    styles = {
        fontSize: '10px',
        fontWeight: 'bold'
    };
    
    render() {
        
        
        return (
            <div>>
                <span className={ this.getBadgeClasses() } style={this.styles}>{ this.formatCount() }</span>
                {/* - adding inline style: */}
                <span style={{ fontSize: 30 }}>Hello</span>
                <button>Increment</button>
                {/*- how to render an array(needs top be set a uniqe key to this list): */}
                <ul> { this.state.tags.map(tag => <li key={ tag }>{ tag }</li>) } </ul>
            </div>
        )
    }
    
    // - rendering classes dynamically
    getBadgeClasses() {
        let classes = "badge-";
        classes += (this.state.count === 0) ? "warning" : "primary"; 
        return classes;
    }
    
    formatCount() {
        // - using object destructuring:
        const { count }= this.state;
        // - jsx tags work like any other react element: 
        const one = <h1>One</h1>; 
        return count === 0 ? <h1>Zero</h1> : count;
    }
}
```

# 2.2 Conditional rendering

```jsx
class Counter extends Component {
    state = {
        count: 1,
        tags: ["tag1", "tag2", "tag3"]
    };
    
    renderTags() {
        if (this.state.tags.length === 0) return <p>'There are no tags'</p>;
        return  <ul> { this.state.tags.map(tag => <li key={ tag }>{ tag }</li>) } </ul>
    }
    
    render() {
        return (
            <div>
                {/* -converting the string intop truthy: */}
                { this.state.tags.length === 0 && 'Please create a new tag!' }
                { this.renderTags() }
            </div>
        )
    }
}
```

# 2.3 Handling events, binding event handler, updating state

```jsx
class Counter extends Component {
    state = {
        count: 1,
    };
    
    // - !without the constructor method we cannot acces "this" in handleIncrement()
    // constructor() {
    //     super();
    //     this.handleIncrement = this.handleIncrement.bind(this);
    // }
    //
    // handleIncrement() {
    //     console.log('Increment clicked', this);
    // }

    // - to access "this" we need to have an arrow function as handleIncrement()
    handleIncrement = () => {
        // this.state.count++; - this is NOT correct!!!, we don't modify state directly!!!
        this.setState({ count: this.state.count + 1 })
    }
    
    render() {
        return (
            <div>
                <span className={ this.getBadgeClasses() } style={this.styles}>{ this.formatCount() }</span>
                {/* - we are not calling the function, we are just simply passing a reference to it: */}
                <button onClick={this.handleIncrement} >Increment</button>
            </div>
        )
    }

    formatCount() {
        // - using object destructuring:
        const { count }= this.state;
        // - jsx tags work like any other react element: 
        const one = <h1>One</h1>;
        return count === 0 ? <h1>Zero</h1> : count;
    }
}
```

# 2.4 Passing event arguments

```jsx
class Counter extends Component {
    state = {
        count: 1,
    };
    
    handleIncrement = (product) => {
        console.log(product);
        this.setState({ count: this.state.count + 1 })
    }
    
    render() {
        return (
            <div>
                <span>{ this.state.count }</span>
                <button onClick={() => this.handleIncrement(product)} >Increment</button>
            </div>
        )
    }
}
```

Ch3 - Composing components
-----------------------
# 3.1 Passing data to Components, Raising and handling events
* We can pass data to a child component via props.
* If we pass something between the tags of a component that can be accessed as props.children.
* Props are input to a component vs state is local to a component.
* Props is read only.
* Modifying a state should be always done in the same component.
* Components can raise events ex. onDelete in the child comp, and handleDelete() in the parent comp where we want to 
* modify the state. Basically we have a method in the parent component, and we pass a reference to the child component 
* via props.

```jsx
// counters file
import Counter from "./counter"
class Counters extends Component {
    state = {
      counters: [
          { id: 1, value: 0 },
          { id: 2, value: 0 },
          { id: 3, value: 0 }
      ]
    };
    
    handleDelete = () => {
        console.log("Event handler called.");
    };
    
    render() {
        return (
            <div>
                { this.state.counters.map(counter => (
                    <Counter key={counter.id} value={counter.value} selected={true} onDelete={this.handleDelete}>
                        {/* - this creates a children props: */}
                        <h4> Counter nr. {counter.id} </h4>
                    </Counter>
                ))}
            </div>);
    }
}

// counter file
class Counter extends Component {
    state = {
        value: this.props.value,
    };
    
    handleIncrement = (product) => {
        console.log(product);
        this.setState({ value: this.state.value + 1 })
    }
    
    render() {
        
        return (
            <div>
                {this.props.children}
                <span>{ this.state.value }</span>
                <button onClick={() => this.handleIncrement(product)} >Increment</button>
                <button onClick={this.props.onDelete} >Delete</button>
            </div>
        )
    }
}
```