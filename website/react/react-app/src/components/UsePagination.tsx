export const UsePagination = ({
    totalCount,
    pageSize,
    siblingCount = 1,
    currentPage
}) => {
    const paginationRange = useMemo(() => {
        // Logic will go here
    }, [totalCount, pageSize, siblingCount, currentPage]);

    return paginationRange;
}