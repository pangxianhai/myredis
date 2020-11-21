SELECT
exp_id as eid
,dp as `time`
,count(distinct device_id) as exp_number
,count(distinct if(product_page_pv>0,device_id,null)) as `商详UV`
from
(
select device_id,dp
from soyoung_dw.dm_inp_abtest_device_all_d
where dp='{before_1}' and exp_id in ({eid_list})
) a
inner JOIN (
    select device_id,dp from soyoung_dw.dwd_md_app_log_nocheat
    where (from_action_id in ({page_action_list}) or curr_page_id in ({page_action_list}))
    and dp='{before_1}'
) t
on (a.device_id = t.device_id and a.dp = t.dp)
group by exp_id,dp
