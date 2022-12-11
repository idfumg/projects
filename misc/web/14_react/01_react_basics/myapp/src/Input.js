import { forwardRef } from "react";

const Input = forwardRef((props, ref) => {
    return (
        <div className='mb-3'>
            <label htmlFor={props.name} className='form-label'>{props.title}</label>
            <input
                type={props.type}
                className={props.className}
                id={props.name}
                autoComplete={props.autoComplete}
                ref={ref}
                onChange={props.onChange}
            ></input>
        </div>
    )
})

export default Input;