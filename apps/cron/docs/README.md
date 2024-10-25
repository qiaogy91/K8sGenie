
#### 说明
1. namespace_record ：每次创建前会清楚当前的重复数据，因此支持一个月创建31次
2. project_record: 每次创建前会清除当月的重复数据，根据当月的 namespace_record 进行一次统计
3. line_record: 每次创建前会清除当月的重复数据，根据当月的 project_record 进行一次统计