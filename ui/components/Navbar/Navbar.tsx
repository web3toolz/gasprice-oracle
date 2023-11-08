import s from './Navbar.module.css';


export default function Navbar() {
    return (
        <nav className={s.root}>
            <div className="max-w-6xl px-6 mx-auto">
                <div className="relative flex flex-row justify-between py-4 align-center md:py-6">
                    <div className="flex items-center flex-1">
                        <a href="/" className="flex items-center">
                            Oracle
                        </a>
                    </div>
                </div>
            </div>
        </nav>
    )
};