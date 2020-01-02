export const validGrid = (grid) => {
    const areaCounts = grid.reduce((acc, row) => {        
        row.forEach(cell => {
            let areaCount = acc[cell.areaId]
            if (!areaCount) {
                acc[cell.areaId] = 1
            } else {
                acc[cell.areaId] += 1
            }
        })
        return acc
    }, {});
    return !Object.values(areaCounts).find(count => count < 4);
}
