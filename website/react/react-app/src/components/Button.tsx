import React from 'react'

interface Properties {
    children: string;
    color?: 'primary' | 'secondary' | 'danger';
    onClick: () => void;
}

const Button = ({ children, onClick, color = 'primary' }: Properties) => {
    return (
        <button className={ "btn btn-" + color } onClick={ onClick }>{ children }</button>
    )
}

export default Button