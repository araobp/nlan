##Do I have to use OpenFlow controller for NLAN?

I would say, "No", because all the OF flows are rather static and do not require dynamic control.

OF flows are ephemeral, but I would rather want a capability to add persistent flows that survives over reboots.

In case of NLAN, all the flows will be resumed by referring to NLAN state on a local database. neutron-lan used ovsdb as a general-purpose local database, but go-nlan may use another data base or a flat file. I have not decided yet.

See also this: http://openvswitch.org/pipermail/dev/2015-January/050380.html

