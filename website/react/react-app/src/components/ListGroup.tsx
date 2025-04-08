function ListGroup() {
    const items = [
        'New York',
        'San Francisco',
        'Tokyo',
        'London',
        'Paris'
    ]

    return (
        <>
            <h1>This is what we might call a saucy list.</h1>
            { items.length === 0 && <p>No items found!!!</p> }
            <ul className="list-group">
                {items.map(item => <li key={item}>{item}</li>)}
            </ul>
        </>
    );
}

export default ListGroup