// import { MouseEvent } from "react";
import { useState } from "react";
interface ListGroupProperties {
    items: string[];
    heading: string;
    // (item: string) => void
    onSelectItem: (item: string) => void;
}

function ListGroup({ items, heading, onSelectItem }: ListGroupProperties) {

    // This is a hook
    const [selectedIndex, setSelectedIndex] = useState(-1);
    // chump_array[0] // variable (selectedIndex)
    // chump_array[1] // updater function

    // Event Handler
    // const handleClick = (event: MouseEvent) => console.log(event)

    return (
        <>
            <h1>{ heading }</h1>
            { items.length === 0 && <p>No items found!!!</p> }
            <ul className="list-group">
                {items.map((item, index) => (
                    <li 
                        className={ selectedIndex === index ? 'list-group-item active' : 'list-group-item' } 
                        key={item} 
                        // onClick={handleClick}
                        onClick={() => { 
                            setSelectedIndex(index); 
                            onSelectItem(item);
                        }}
                    >
                        {item}
                    </li>
                ))}
            </ul>
        </>
    );
}

export default ListGroup