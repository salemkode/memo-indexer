import Page from "../../components/page";
import Pagination from "../../components/util/pagination";
import {useEffect, useState} from "react";
import Link from 'next/link';
import styles from '../../styles/list.module.css';
import dropdownStyles from '../../styles/dropdown.module.css';

function List() {
    const [loading, setLoading] = useState(true)
    const [allPeers, setAllPeers] = useState([])
    const [peers, setPeers] = useState([])
    const [errorMessage, setErrorMessage] = useState("")
    const [totalPeers, setTotalPeers] = useState(0)
    const [selectValue, setSelectValue] = useState("attempted")

    useEffect(() => {
        fetch("/api/peers").then(res => {
            if (res.ok) {
                return res.json()
            }
            return Promise.reject(res)
        }).then(data => {
            setAllPeers(data.Peers)
            setTotalPeers(data.Peers.length)
            setLoading(false)
        }).catch(res => {
            res.text().then(msg => {
                setErrorMessage(<>Code: {res.status}<br/>Message: {msg}</>)
            })
        })
    }, [])

    const onPageChanged = (data) => {
        const {currentPage, totalPages, pageLimit} = data
        const offset = (currentPage - 1) * pageLimit
        setPeers(allPeers.slice(offset, offset + pageLimit))
    }

    const selectOnChange = (event) => {
        console.log(event.target.value)
        setSelectValue(event.target.value)
    }

    return (
        <Page>
            <div>
                <h1>
                    Peers Page
                </h1>
                <div>
                    <select className={dropdownStyles.select} onChange={selectOnChange} value={selectValue}>
                        <option value={"all"}>All</option>
                        <option value={"attempted"}>Attempted</option>
                        <option value={"successes"}>Successes</option>
                    </select>
                </div>
                {loading ?
                    <>{!!errorMessage ?
                        <>Error: {errorMessage}</>
                        :
                        <>Loading...</>
                    }</>
                    :
                    <div>
                        <ul className={styles.list}>
                            {peers.map((peer, key) => (
                                <li key={key}>
                                    <Link href={{
                                        pathname: "/peer/view",
                                        query: {
                                            ip: peer.Ip,
                                            port: peer.Port
                                        }
                                    }}>
                                        <a>{peer.Ip}:{peer.Port} - {peer.Time} - {peer.Status}</a>
                                    </Link>
                                </li>
                            ))}
                        </ul>
                        <Pagination totalRecords={totalPeers} pageLimit={10} pageNeighbours={1}
                                    onPageChanged={onPageChanged}/>
                    </div>
                }
            </div>
        </Page>
    )
}

export default List
