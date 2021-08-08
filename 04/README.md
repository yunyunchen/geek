new Application Project Layout
- /api 
  - APi协议定义目录
- /cmd
  - 整个项目启动的入口文件
- /configs
    - 配置文件模板或默认配置
- /internal
  - biz 业务逻辑的组装层（repo）
  - data 业务数据访问层（原model），cache，db封装
  - service api定义的服务层，处理DTO到biz领域实体的转换；协同各类 biz 交互，但是不应处理复杂逻辑。
    - ps：（gRPC的message到,它到一个biz的领域实体的一个转化。也就是说，它需要把DTO对象转换成一个domain object，domain object定义在biz层。biz层最终暴露的是，service层拿到的，gRPC的DTO对象，那么你去调用biz层的时候，一定要做一个deep copy，拿到一个domain object，才能用这个方法去调用，往下传递。做一些简单的逻辑层的领域对象的一些组装。没有复杂逻辑，只是编排各种对象，来完成一个复杂的业务场景的。
  
DTO ,数据传输对象

- internal: 是为了避免有同业务下有人跨目录引用了内部的 biz、data、service 等内部 struct。
    - biz: 业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo，repo 接口在这里定义，使用依赖倒置的原则。
    - data: 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。我们可能会把 data 与 dao 混淆在一起，data 偏重业务的含义，它所要做的是将领域对象重新拿出来，我们去掉了 DDD 的 infra 层。
    - service: 实现了 api 定义的服务层，类似 DDD 的 application 层，处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑。
    - PO(Persistent Object): 持久化对象，它跟持久层（通常是关系型数据库）的数据结构形成一一对应的映射关系，如果持久层是关系型数据库，那么数据表中的每个字段（或若干个）就对应 PO 的一个（或若干个）属性。https://github.com/facebook/ent


