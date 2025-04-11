

function Message() {

    const name = 'Ryan'

    // This is JSX - Javascript XML
    if (name)
        return <h1>Adunarino, {name}</h1>;
    return <h1>Adunarino, Torias</h1>;
}

export default Message;