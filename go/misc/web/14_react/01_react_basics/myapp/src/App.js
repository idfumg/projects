import React, { Fragment, useEffect, useRef, useState } from 'react';
import './App.css';
import Input from './Input';

function App(props) {
    const [isTrue, setIsTrue] = useState(true);
    const [crowd, setCrowd] = useState([]);
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [dob, setDob] = useState("");

    // refs
    const firstNameRef = useRef();
    const lastNameRef = useRef(null);
    const dobRef = useRef(null);

    const toggleTrue = () => {
        if (isTrue) {
            setIsTrue(false);
        } else {
            setIsTrue(true);
        }
    };

    useEffect(() => {
        console.log("useEffect fired!")

        let people = [
            {
                id: 1,
                firstName: "Mary",
                lastName: "Jones",
                dob: "1997-05-02",
            },
            {
                id: 2,
                firstName: "Jack",
                lastName: "Smith",
                dob: "1999-02-04",
            },
        ]

        setCrowd(people);
    }, []);

    const handleSubmit = (event) => {
        event.preventDefault();
        if (lastName !== "") {
            addPerson(firstName, lastName, dob);
        }
    }

    const addPerson = (newFirstName, newLastName, newDob) => {
        let newPerson = {
            id: crowd.length + 1,
            firstName: newFirstName,
            lastName: newLastName,
            dob: newDob,
        }

        const newList = crowd.concat(newPerson)

        const sortedList = newList.sort((a, b) => {
            if (a.lastName < b.lastName) {
                return -1;
            } else if (a.lastName > b.lastName) {
                return 1;
            } else {
                return 0;
            }
        })

        setCrowd(sortedList);
        setFirstName("");
        setLastName("");
        setDob("");
        firstNameRef.current.value = ""
        lastNameRef.current.value = ""
        dobRef.current.value = ""
    }

    return (
        <Fragment>
            <hr />
            <h1 className="h1-green">{props.msg}</h1>
            <hr />
            {isTrue &&
                <Fragment>
                    <p>The current value of isTrue is true</p>
                    <hr></hr>
                </Fragment>
            }
            <hr />
            {isTrue ? <p>Is true</p> : <p>Is false</p>}
            <hr />
            <a href="#!" className="btn btn-outline-secondary" onClick={toggleTrue}>Toggle isTrue</a>
            <hr />
            <form autoComplete='off' onSubmit={handleSubmit}>
                <div className='mb-3'>
                    <label className='form-label' htmlFor='first-name'>First name</label>
                    <input 
                        className='form-control' 
                        id='first-name' 
                        type='text' 
                        name='first-name' 
                        autoComplete='first-name-new'
                        ref={firstNameRef}
                        onChange={(event) => setFirstName(event.target.value)}
                    ></input>
                </div>

                <Input
                    title='Last Name'
                    type='text'
                    name='last-name'
                    autoComplete='last-name-new'
                    className='form-control'
                    ref={lastNameRef}
                    onChange={(event) => setLastName(event.target.value)}
                ></Input>

                <Input
                    title='Birthday'
                    type='date'
                    name='dob'
                    autoComplete='dob-new'
                    className='form-control'
                    ref={dobRef}
                    onChange={(event) => setDob(event.target.value)}
                ></Input>

                <input type='submit' value='Submit' className='btn btn-primary'></input>
            </form>

            <div>
                First Name: {firstName}<br />
                Last Name: {lastName}<br />
                Birthday: {dob}<br />
            </div>

            <hr />
            <h3>People</h3>
            <ul className="list-group">
                {crowd.map((m) => (
                    <li key={m.id} className="list-group-item">{m.firstName} {m.lastName}</li>
                ))}
            </ul>
        </Fragment>
    )
}

export default App;
