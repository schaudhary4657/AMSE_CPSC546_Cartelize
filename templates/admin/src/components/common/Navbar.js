import React from 'react'
import {
    Navbar,
    NavbarBrand,
} from 'reactstrap'

export default function Header(props) {
    return (
        <Navbar color="primary">
            <NavbarBrand className="text-light">Cartelize</NavbarBrand>
        </Navbar>
    )
}