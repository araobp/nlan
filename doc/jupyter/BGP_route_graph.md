

```python
import networkx as nx
get_ipython().magic('matplotlib inline')
```


```python
import tega.driver
d = tega.driver.Driver(host='192.168.57.133')
subnets = d.get(path='graph.subnets')
```


```python
g = nx.DiGraph(subnets['172.21.1.0/24'])
nx.draw_spring(g, node_size=1000, with_labels=True, arrows=True, alpha=0.8)
```


![png](output_2_0.png)

