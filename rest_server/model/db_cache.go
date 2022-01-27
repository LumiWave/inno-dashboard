package model

func (o *DB) HKeys(key string) ([]string, error) {
	return o.Cache.HKeys(key)
}
