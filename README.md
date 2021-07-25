# geektime_homework

## 第二周
### 问题
+ 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

### 回答
+ 应该,dao 层直接将错误上抛，由业务层判断 sql.ErrNoRows 是有可以进行降级处理

```go
func (m *DBModel) queryDepPeople(q string, d int) ([]*DepPeople, error) {
    depPeoples := []*DepPeople{}
    query := fmt.Sprintf(q, d)

    err := m.DBEngine.QueryRow(query).Scan(&depPeoples)
    if err != nil {
        return nil, fmt.Errorf("query sql error: %s,%w", q, err)
    }

    return depPeoples, nil
}


func (m *DBModel) GetAllDepPeople(d int) ([]*DepPeople, error) {
    query := "SELECT  ""
    allPeople, err := m.queryDepPeople(query, d)
    if err != nil && !errors.Is(err, sql.ErrNoRows) {
        return nil, fmt.Errorf("获取部门所有员工失败: %w", err)
    }
    return allPeople, nil
}
```
