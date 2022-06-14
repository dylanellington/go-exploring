<h2>200. Number of Islands</h2>
<h3 style="color:Yellow;">Medium</h3>
<hr>
<div><p>Given an <code>m x n</code> 2D binary grid <code>grid</code> which represents a map of <code>&#39;1&#39;</code>s (land) and <code>&#39;0&#39;</code>s (water), return <em>the number of islands</em>.</p>

<p>An <strong>island</strong> is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.</p>

<p> </p>
<p><strong>Example 1:</strong></p>

<pre><strong>Input:</strong> grid = [
  [&#34;1&#34;,&#34;1&#34;,&#34;1&#34;,&#34;1&#34;,&#34;0&#34;],
  [&#34;1&#34;,&#34;1&#34;,&#34;0&#34;,&#34;1&#34;,&#34;0&#34;],
  [&#34;1&#34;,&#34;1&#34;,&#34;0&#34;,&#34;0&#34;,&#34;0&#34;],
  [&#34;0&#34;,&#34;0&#34;,&#34;0&#34;,&#34;0&#34;,&#34;0&#34;]
]
<strong>Output:</strong> 1
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> grid = [
  [&#34;1&#34;,&#34;1&#34;,&#34;0&#34;,&#34;0&#34;,&#34;0&#34;],
  [&#34;1&#34;,&#34;1&#34;,&#34;0&#34;,&#34;0&#34;,&#34;0&#34;],
  [&#34;0&#34;,&#34;0&#34;,&#34;1&#34;,&#34;0&#34;,&#34;0&#34;],
  [&#34;0&#34;,&#34;0&#34;,&#34;0&#34;,&#34;1&#34;,&#34;1&#34;]
]
<strong>Output:</strong> 3
</pre>

<p> </p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>m == grid.length</code></li>
	<li><code>n == grid[i].length</code></li>
	<li><code>1 &lt;= m, n &lt;= 300</code></li>
	<li><code>grid[i][j]</code> is <code>&#39;0&#39;</code> or <code>&#39;1&#39;</code>.</li>
</ul>
</div>