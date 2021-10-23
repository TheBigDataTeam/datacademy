import { useEffect } from 'react'


export const useDocTitle = (params: string): void => {
    return useEffect(() => {
        document.title = params
    }, [params])

}
