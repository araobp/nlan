

```python
import networkx as nx
get_ipython().magic('matplotlib inline')
```


```python
import tega.driver
d = tega.driver.Driver(host='192.168.57.133')
server_client = d.get(path='graph.server_client.8888')
```


```python
g = nx.DiGraph(server_client)
nx.draw_spring(g, node_size=1000, with_labels=True, arrows=True, alpha=0.8)
```


![png](output_3_0.png)

