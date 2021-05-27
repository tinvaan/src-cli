func TestTaskCacheKey(t *testing.T) {
	key := TaskCacheKey{&Task{
		Repository: testRepo1,
		Steps:      steps,
	}}
func TestExecutionDiskCache_GetSet(t *testing.T) {
	cacheKey1 := TaskCacheKey{Task: &Task{
		Repository: testRepo1,
	cacheKey2 := TaskCacheKey{Task: &Task{
		Repository: testRepo2,
	cache := ExecutionDiskCache{Dir: cacheTmpDir(t)}
	// Empty cache, no hits
	assertCacheMiss(t, cache, cacheKey1)
	assertCacheMiss(t, cache, cacheKey2)
	// Set the cache
	if err := cache.Set(ctx, cacheKey1, value); err != nil {
		t.Fatalf("cache.Set returned unexpected error: %s", err)
	}
	// Cache hit
	assertCacheHit(t, cache, cacheKey1, value)
	// Cache miss due to different key
	assertCacheMiss(t, cache, cacheKey2)
	// Cache miss due to cleared cache
	if err := cache.Clear(ctx, cacheKey1); err != nil {
		t.Fatalf("cache.Get returned unexpected error: %s", err)
	assertCacheMiss(t, cache, cacheKey1)
func assertCacheHit(t *testing.T, c ExecutionDiskCache, k TaskCacheKey, want executionResult) {
func assertCacheMiss(t *testing.T, c ExecutionDiskCache, k TaskCacheKey) {