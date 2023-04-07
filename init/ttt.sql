create table if not exists isc_demo.service_update(
        id bigint(20) not null auto_increment comment '主键',
        service_name varchar(32) default '' not null comment '服务名',
        times varchar(32) default '' not null comment '更新时间',
        primary key (id), 
        unique key uk_service(service_name)
        )engine=InnoDB default charset =utf8mb4 comment='表';