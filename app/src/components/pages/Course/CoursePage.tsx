import React, { useEffect, useState } from 'react';
import axios from 'axios';

export const CoursePage: React.FunctionComponent = (): JSX.Element => {

    const [data, setData] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            const result = await axios.get("http://localhost:3100/courses/1")
            setData(result)
        }
        fetchData();
    }, [])

    return data ? (<h2>{data.data.author}</h2>) : (<h2>Loading</h2>)
}