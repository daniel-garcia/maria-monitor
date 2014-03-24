package mariamon

import (
	"database/sql"
	"errors"
)

type pStatus struct {
	AbortedConnects                       int64 `json:"Aborted_connects"`
	BinlogBytesWritten                    int64 `json:"Binlog_bytes_written"`
	BinlogCacheDiskUse                    int64 `json:"Binlog_cache_disk_use"`
	BinlogCommits                         int64 `json:"Binlog_commits"`
	BinlogGroupCommits                    int64 `json:"Binlog_group_commits"`
	BinlogSnapshotFiles                   int64 `json:"Binlog_snapshot_files"`
	BinlogSnapshotPosition                int64 `json:"Binlog_snapshot_position"`
	BinlogStmtCacheDiskUse                int64 `json:"Binlog_stmt_cache_disk_use"`
	BinlogStmtCacheUse                    int64 `json:"Binlog_stmt_cache_use"`
	BusyTime                              int64 `json:"Busy_time"`
	BytesReceived                         int64 `json:"Bytes_received"`
	ComAdminCommands                      int64 `json:"Com_admin_commands"`
	ComAlterDbUpgrade                     int64 `json:"Com_alter_db_upgrade"`
	ComAlterFunction                      int64 `json:"Com_alter_function"`
	ComAlterServer                        int64 `json:"Com_alter_server"`
	ComAlterTablespace                    int64 `json:"Com_alter_tablespace"`
	ComAssignToKeycache                   int64 `json:"Com_assign_to_keycache"`
	ComBegin                              int64 `json:"Com_begin"`
	ComCallProcedure                      int64 `json:"Com_call_procedure"`
	ComCheck                              int64 `json:"Com_check"`
	ComCommit                             int64 `json:"Com_commit"`
	ComCreateEvent                        int64 `json:"Com_create_event"`
	ComCreateIndex                        int64 `json:"Com_create_index"`
	ComCreateRole                         int64 `json:"Com_create_role"`
	ComCreateServer                       int64 `json:"Com_create_server"`
	ComCreateTrigger                      int64 `json:"Com_create_trigger"`
	ComCreateUser                         int64 `json:"Com_create_user"`
	ComDeallocSql                         int64 `json:"Com_dealloc_sql"`
	ComDeleteMulti                        int64 `json:"Com_delete_multi"`
	ComDropDb                             int64 `json:"Com_drop_db"`
	ComDropFunction                       int64 `json:"Com_drop_function"`
	ComDropProcedure                      int64 `json:"Com_drop_procedure"`
	ComDropServer                         int64 `json:"Com_drop_server"`
	ComDropTrigger                        int64 `json:"Com_drop_trigger"`
	ComDropView                           int64 `json:"Com_drop_view"`
	ComExecuteSql                         int64 `json:"Com_execute_sql"`
	ComGetDiagnostics                     int64 `json:"Com_get_diagnostics"`
	ComGrant                              int64 `json:"Com_grant"`
	ComHaClose                            int64 `json:"Com_ha_close"`
	ComHaRead                             int64 `json:"Com_ha_read"`
	ComInsert                             int64 `json:"Com_insert"`
	ComInstallPlugin                      int64 `json:"Com_install_plugin"`
	ComLoad                               int64 `json:"Com_load"`
	ComLoadMasterTable                    int64 `json:"Com_load_master_table"`
	ComLockTables                         int64 `json:"Com_lock_tables"`
	ComPreloadKeys                        int64 `json:"Com_preload_keys"`
	ComPurge                              int64 `json:"Com_purge"`
	ComReleaseSavepoint                   int64 `json:"Com_release_savepoint"`
	ComRenameUser                         int64 `json:"Com_rename_user"`
	ComReplace                            int64 `json:"Com_replace"`
	ComReset                              int64 `json:"Com_reset"`
	ComRestoreTable                       int64 `json:"Com_restore_table"`
	ComRevoke                             int64 `json:"Com_revoke"`
	ComRevokeGrant                        int64 `json:"Com_revoke_grant"`
	ComRollback                           int64 `json:"Com_rollback"`
	ComSavepoint                          int64 `json:"Com_savepoint"`
	ComSetOption                          int64 `json:"Com_set_option"`
	ComShowAuthors                        int64 `json:"Com_show_authors"`
	ComShowBinlogs                        int64 `json:"Com_show_binlogs"`
	ComShowClientStatistics               int64 `json:"Com_show_client_statistics"`
	ComShowCollations                     int64 `json:"Com_show_collations"`
	ComShowContributors                   int64 `json:"Com_show_contributors"`
	ComShowCreateEvent                    int64 `json:"Com_show_create_event"`
	ComShowCreateProc                     int64 `json:"Com_show_create_proc"`
	ComShowCreateTrigger                  int64 `json:"Com_show_create_trigger"`
	ComShowEngineLogs                     int64 `json:"Com_show_engine_logs"`
	ComShowEngineStatus                   int64 `json:"Com_show_engine_status"`
	ComShowErrors                         int64 `json:"Com_show_errors"`
	ComShowFields                         int64 `json:"Com_show_fields"`
	ComShowGrants                         int64 `json:"Com_show_grants"`
	ComShowIndexStatistics                int64 `json:"Com_show_index_statistics"`
	ComShowOpenTables                     int64 `json:"Com_show_open_tables"`
	ComShowPrivileges                     int64 `json:"Com_show_privileges"`
	ComShowProcesslist                    int64 `json:"Com_show_processlist"`
	ComShowProfiles                       int64 `json:"Com_show_profiles"`
	ComShowStatus                         int64 `json:"Com_show_status"`
	ComShowTableStatistics                int64 `json:"Com_show_table_statistics"`
	ComShowTables                         int64 `json:"Com_show_tables"`
	ComShowUserStatistics                 int64 `json:"Com_show_user_statistics"`
	ComShowWarnings                       int64 `json:"Com_show_warnings"`
	ComStmtClose                          int64 `json:"Com_stmt_close"`
	ComStmtFetch                          int64 `json:"Com_stmt_fetch"`
	ComStmtReprepare                      int64 `json:"Com_stmt_reprepare"`
	ComStmtSendLongData                   int64 `json:"Com_stmt_send_long_data"`
	ComUninstallPlugin                    int64 `json:"Com_uninstall_plugin"`
	ComUpdate                             int64 `json:"Com_update"`
	ComXaCommit                           int64 `json:"Com_xa_commit"`
	ComXaPrepare                          int64 `json:"Com_xa_prepare"`
	ComXaRollback                         int64 `json:"Com_xa_rollback"`
	Compression                           int64 `json:"Compression"`
	ConnectionErrorsInternal              int64 `json:"Connection_errors_internal"`
	ConnectionErrorsMaxConnections        int64 `json:"Connection_errors_max_connections"`
	ConnectionErrorsPeerAddress           int64 `json:"Connection_errors_peer_address"`
	ConnectionErrorsSelect                int64 `json:"Connection_errors_select"`
	ConnectionErrorsTcpwrap               int64 `json:"Connection_errors_tcpwrap"`
	Connections                           int64 `json:"Connections"`
	CreatedTmpDiskTables                  int64 `json:"Created_tmp_disk_tables"`
	CreatedTmpTables                      int64 `json:"Created_tmp_tables"`
	DelayedInsertThreads                  int64 `json:"Delayed_insert_threads"`
	EmptyQueries                          int64 `json:"Empty_queries"`
	ExecutedEvents                        int64 `json:"Executed_events"`
	ExecutedTriggers                      int64 `json:"Executed_triggers"`
	FeatureDynamicColumns                 int64 `json:"Feature_dynamic_columns"`
	FeatureFulltext                       int64 `json:"Feature_fulltext"`
	FeatureGis                            int64 `json:"Feature_gis"`
	FeatureLocale                         int64 `json:"Feature_locale"`
	FeatureSubquery                       int64 `json:"Feature_subquery"`
	FeatureTimezone                       int64 `json:"Feature_timezone"`
	FeatureTrigger                        int64 `json:"Feature_trigger"`
	FeatureXml                            int64 `json:"Feature_xml"`
	FlushCommands                         int64 `json:"Flush_commands"`
	HandlerDelete                         int64 `json:"Handler_delete"`
	HandlerExternalLock                   int64 `json:"Handler_external_lock"`
	HandlerIcpAttempts                    int64 `json:"Handler_icp_attempts"`
	HandlerIcpMatch                       int64 `json:"Handler_icp_match"`
	HandlerMrrInit                        int64 `json:"Handler_mrr_init"`
	HandlerMrrKeyRefills                  int64 `json:"Handler_mrr_key_refills"`
	HandlerMrrRowidRefills                int64 `json:"Handler_mrr_rowid_refills"`
	HandlerPrepare                        int64 `json:"Handler_prepare"`
	HandlerReadKey                        int64 `json:"Handler_read_key"`
	HandlerReadNext                       int64 `json:"Handler_read_next"`
	HandlerReadRnd                        int64 `json:"Handler_read_rnd"`
	HandlerReadRndNext                    int64 `json:"Handler_read_rnd_next"`
	HandlerSavepoint                      int64 `json:"Handler_savepoint"`
	HandlerTmpUpdate                      int64 `json:"Handler_tmp_update"`
	HandlerTmpWrite                       int64 `json:"Handler_tmp_write"`
	HandlerUpdate                         int64 `json:"Handler_update"`
	KeyBlocksNotFlushed                   int64 `json:"Key_blocks_not_flushed"`
	KeyBlocksUsed                         int64 `json:"Key_blocks_used"`
	KeyReadRequests                       int64 `json:"Key_read_requests"`
	KeyWriteRequests                      int64 `json:"Key_write_requests"`
	LastQueryCost                         int64 `json:"Last_query_cost"`
	MaxUsedConnections                    int64 `json:"Max_used_connections"`
	NotFlushedDelayedRows                 int64 `json:"Not_flushed_delayed_rows"`
	OpenStreams                           int64 `json:"Open_streams"`
	OpenTables                            int64 `json:"Open_tables"`
	OpenedTableDefinitions                int64 `json:"Opened_table_definitions"`
	OpenedViews                           int64 `json:"Opened_views"`
	OqgraphBoostVersion                   int64 `json:"Oqgraph_boost_version"`
	PerformanceSchemaAccountsLost         int64 `json:"Performance_schema_accounts_lost"`
	PerformanceSchemaCondClassesLost      int64 `json:"Performance_schema_cond_classes_lost"`
	PerformanceSchemaCondInstancesLost    int64 `json:"Performance_schema_cond_instances_lost"`
	PerformanceSchemaDigestLost           int64 `json:"Performance_schema_digest_lost"`
	PerformanceSchemaFileClassesLost      int64 `json:"Performance_schema_file_classes_lost"`
	PerformanceSchemaFileHandlesLost      int64 `json:"Performance_schema_file_handles_lost"`
	PerformanceSchemaFileInstancesLost    int64 `json:"Performance_schema_file_instances_lost"`
	PerformanceSchemaHostsLost            int64 `json:"Performance_schema_hosts_lost"`
	PerformanceSchemaLockerLost           int64 `json:"Performance_schema_locker_lost"`
	PerformanceSchemaMutexClassesLost     int64 `json:"Performance_schema_mutex_classes_lost"`
	PerformanceSchemaMutexInstancesLost   int64 `json:"Performance_schema_mutex_instances_lost"`
	PerformanceSchemaRwlockClassesLost    int64 `json:"Performance_schema_rwlock_classes_lost"`
	PerformanceSchemaRwlockInstancesLost  int64 `json:"Performance_schema_rwlock_instances_lost"`
	PerformanceSchemaSocketClassesLost    int64 `json:"Performance_schema_socket_classes_lost"`
	PerformanceSchemaSocketInstancesLost  int64 `json:"Performance_schema_socket_instances_lost"`
	PerformanceSchemaStageClassesLost     int64 `json:"Performance_schema_stage_classes_lost"`
	PerformanceSchemaStatementClassesLost int64 `json:"Performance_schema_statement_classes_lost"`
	PerformanceSchemaTableHandlesLost     int64 `json:"Performance_schema_table_handles_lost"`
	PerformanceSchemaTableInstancesLost   int64 `json:"Performance_schema_table_instances_lost"`
	PerformanceSchemaThreadClassesLost    int64 `json:"Performance_schema_thread_classes_lost"`
	PerformanceSchemaThreadInstancesLost  int64 `json:"Performance_schema_thread_instances_lost"`
	PerformanceSchemaUsersLost            int64 `json:"Performance_schema_users_lost"`
	PreparedStmtCount                     int64 `json:"Prepared_stmt_count"`
	QcacheFreeBlocks                      int64 `json:"Qcache_free_blocks"`
	QcacheHits                            int64 `json:"Qcache_hits"`
	QcacheLowmemPrunes                    int64 `json:"Qcache_lowmem_prunes"`
	QcacheQueriesInCache                  int64 `json:"Qcache_queries_in_cache"`
	Queries                               int64 `json:"Queries"`
	RowsRead                              int64 `json:"Rows_read"`
	RowsSent                              int64 `json:"Rows_sent"`
	RowsTmpRead                           int64 `json:"Rows_tmp_read"`
	RplRecoveryRank                       int64 `json:"Rpl_recovery_rank"`
	SelectFullRangeJoin                   int64 `json:"Select_full_range_join"`
	SelectRangeCheck                      int64 `json:"Select_range_check"`
	SlowLaunchThreads                     int64 `json:"Slow_launch_threads"`
	SortMergePasses                       int64 `json:"Sort_merge_passes"`
	SortRows                              int64 `json:"Sort_rows"`
	SphinxError                           int64 `json:"Sphinx_error"`
	SphinxTime                            int64 `json:"Sphinx_time"`
	SphinxTotal                           int64 `json:"Sphinx_total"`
	SphinxTotalFound                      int64 `json:"Sphinx_total_found"`
	SphinxWordCount                       int64 `json:"Sphinx_word_count"`
	SphinxWords                           int64 `json:"Sphinx_words"`
	SslAcceptRenegotiates                 int64 `json:"Ssl_accept_renegotiates"`
	SslCallbackCacheHits                  int64 `json:"Ssl_callback_cache_hits"`
	SslCipherList                         int64 `json:"Ssl_cipher_list"`
	SslConnectRenegotiates                int64 `json:"Ssl_connect_renegotiates"`
	SslCtxVerifyMode                      int64 `json:"Ssl_ctx_verify_mode"`
	SslFinishedAccepts                    int64 `json:"Ssl_finished_accepts"`
	SslServerNotAfter                     int64 `json:"Ssl_server_not_after"`
	SslSessionCacheHits                   int64 `json:"Ssl_session_cache_hits"`
	SslSessionCacheMode                   int64 `json:"Ssl_session_cache_mode"`
	SslSessionCacheSize                   int64 `json:"Ssl_session_cache_size"`
	SslSessionsReused                     int64 `json:"Ssl_sessions_reused"`
	SslVerifyDepth                        int64 `json:"Ssl_verify_depth"`
	SslVersion                            int64 `json:"Ssl_version"`
	SubqueryCacheMiss                     int64 `json:"Subquery_cache_miss"`
	Syncs                                 int64 `json:"Syncs"`
	TableLocksImmediate                   int64 `json:"Table_locks_immediate"`
	TcLogMaxPagesUsed                     int64 `json:"Tc_log_max_pages_used"`
	TcLogPageWaits                        int64 `json:"Tc_log_page_waits"`
	ThreadpoolThreads                     int64 `json:"Threadpool_threads"`
	ThreadsCached                         int64 `json:"Threads_cached"`
	ThreadsCreated                        int64 `json:"Threads_created"`
	Uptime                                int64 `json:"Uptime"`
}
type Status struct {
	pStatus
}

func (status Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(status)
}
func (status Status) UnmarshalJSON(blob []byte) error {
	return json.Unmarshal(blob, &status.pStatus)
}

// AbortedConnects returns number of failed server connection attempts. This can be due to a client using an incorrect password, a client not having privileges to connect to a database, a connection packet not containing the correct information, or if it takes more than connect_timeout https://mariadb.com/kb/en/server-system-variables/#connect_timeout seconds to get a connect packet.
func (s *Status) AbortedConnects() int64 {
	return s.pStatus.abortedConnects // Aborted_connects
}

// BinlogBytesWritten returns the number of bytes written to the binary log https://mariadb.com/kb/en/binary-log/.
func (s *Status) BinlogBytesWritten() int64 {
	return s.pStatus.binlogBytesWritten // Binlog_bytes_written
}

// BinlogCacheDiskUse returns number of transactions which used a temporary disk cache because they could not fit in the regular binary log https://mariadb.com/kb/en/binary-log/ cache, being larger than binlog_cache_size </kb/en/server-system-variables/#binlog_cache_size>.
func (s *Status) BinlogCacheDiskUse() int64 {
	return s.pStatus.binlogCacheDiskUse // Binlog_cache_disk_use
}

// BinlogCommits returns total number of transactions committed to the binary log.
func (s *Status) BinlogCommits() int64 {
	return s.pStatus.binlogCommits // Binlog_commits
}

// BinlogGroupCommits returns total number of group commits done to the binary log. See Group commit for the binary log https://mariadb.com/kb/en/group-commit-for-the-binary-log/.
func (s *Status) BinlogGroupCommits() int64 {
	return s.pStatus.binlogGroupCommits // Binlog_group_commits
}

// BinlogSnapshotFiles returns the binary log file. Unlike SHOW MASTER STATUS https://mariadb.com/kb/en/show-master-status/, can be queried in a transactionally consistent way, irrespective of which other transactions have been committed since the snapshot was taken. See Enhancements for START TRANSACTION WITH CONSISTENT SNAPSHOT </kb/en/enhancements-for-start-transaction-with-consistent-snapshot/>.
func (s *Status) BinlogSnapshotFiles() int64 {
	return s.pStatus.binlogSnapshotFiles // Binlog_snapshot_files
}

// BinlogSnapshotPosition returns the binary log position. Unlike SHOW MASTER STATUS https://mariadb.com/kb/en/show-master-status/, can be queried in a transactionally consistent way, irrespective of which other transactions have been committed since the snapshot was taken. See Enhancements for START TRANSACTION WITH CONSISTENT SNAPSHOT </kb/en/enhancements-for-start-transaction-with-consistent-snapshot/>.
func (s *Status) BinlogSnapshotPosition() int64 {
	return s.pStatus.binlogSnapshotPosition // Binlog_snapshot_position
}

// BinlogStmtCacheDiskUse returns number of non-transaction statements which used a temporary disk cache because they could not fit in the regular binary log https://mariadb.com/kb/en/binary-log/ cache, being larger than binlog_stmt_cache_size </kb/en/server-system-variables/#binlog_stmt_cache_size>.
func (s *Status) BinlogStmtCacheDiskUse() int64 {
	return s.pStatus.binlogStmtCacheDiskUse // Binlog_stmt_cache_disk_use
}

// BinlogStmtCacheUse returns number of non-transaction statement which used the regular binary log https://mariadb.com/kb/en/binary-log/ cache, being smaller than binlog_stmt_cache_size </kb/en/server-system-variables/#binlog_stmt_cache_size>.
func (s *Status) BinlogStmtCacheUse() int64 {
	return s.pStatus.binlogStmtCacheUse // Binlog_stmt_cache_use
}

// BusyTime returns cumulative time in seconds of activity on connections.
func (s *Status) BusyTime() int64 {
	return s.pStatus.busyTime // Busy_time
}

// BytesReceived returns total bytes received from all clients.
func (s *Status) BytesReceived() int64 {
	return s.pStatus.bytesReceived // Bytes_received
}

// ComAdminCommands returns number of admin commands executed. These include table dumps, change users, binary log dumps, shutdowns, pings and debugs.
func (s *Status) ComAdminCommands() int64 {
	return s.pStatus.comAdminCommands // Com_admin_commands
}

// ComAlterDbUpgrade returns number of ALTER DATABASE ... UPGRADE https://mariadb.com/kb/en/alter-database/ commands executed.
func (s *Status) ComAlterDbUpgrade() int64 {
	return s.pStatus.comAlterDbUpgrade // Com_alter_db_upgrade
}

// ComAlterFunction returns number of ALTER FUNCTION https://mariadb.com/kb/en/alter-function/ commands executed.
func (s *Status) ComAlterFunction() int64 {
	return s.pStatus.comAlterFunction // Com_alter_function
}

// ComAlterServer returns number of ALTER SERVER https://mariadb.com/kb/en/alter-server/ commands executed.
func (s *Status) ComAlterServer() int64 {
	return s.pStatus.comAlterServer // Com_alter_server
}

// ComAlterTablespace returns number of ALTER TABLESPACE https://mariadb.com/kb/en/alter-tablespace/ commands executed.
func (s *Status) ComAlterTablespace() int64 {
	return s.pStatus.comAlterTablespace // Com_alter_tablespace
}

// ComAssignToKeycache returns number of assign to keycache commands executed.
func (s *Status) ComAssignToKeycache() int64 {
	return s.pStatus.comAssignToKeycache // Com_assign_to_keycache
}

// ComBegin returns number of BEGIN https://mariadb.com/kb/en/begin-end/ statements executed.
func (s *Status) ComBegin() int64 {
	return s.pStatus.comBegin // Com_begin
}

// ComCallProcedure returns number of CALL https://mariadb.com/kb/en/call/ procedure_name statements executed.
func (s *Status) ComCallProcedure() int64 {
	return s.pStatus.comCallProcedure // Com_call_procedure
}

// ComCheck returns number of CHECK TABLE https://mariadb.com/kb/en/check-table/ commands executed.
func (s *Status) ComCheck() int64 {
	return s.pStatus.comCheck // Com_check
}

// ComCommit returns number of COMMIT https://mariadb.com/kb/en/transactions-commit-statement/ commands executed. Differs from Handler_commit <#handler_commit>, which counts internal commit statements.
func (s *Status) ComCommit() int64 {
	return s.pStatus.comCommit // Com_commit
}

// ComCreateEvent returns number of CREATE EVENT https://mariadb.com/kb/en/create-event/ commands executed. Differs from Executed_events <#executed_events> in that it is incremented when the CREATE EVENT is run, and not when the event executes.
func (s *Status) ComCreateEvent() int64 {
	return s.pStatus.comCreateEvent // Com_create_event
}

// ComCreateIndex returns number of CREATE INDEX https://mariadb.com/kb/en/create-index/ commands executed.
func (s *Status) ComCreateIndex() int64 {
	return s.pStatus.comCreateIndex // Com_create_index
}

// ComCreateRole returns number of CREATE ROLE https://mariadb.com/kb/en/create-role/ commands executed.
func (s *Status) ComCreateRole() int64 {
	return s.pStatus.comCreateRole // Com_create_role
}

// ComCreateServer returns number of CREATE SERVER https://mariadb.com/kb/en/create-server/ commands executed.
func (s *Status) ComCreateServer() int64 {
	return s.pStatus.comCreateServer // Com_create_server
}

// ComCreateTrigger returns number of CREATE TRIGGER https://mariadb.com/kb/en/create-trigger/ commands executed.
func (s *Status) ComCreateTrigger() int64 {
	return s.pStatus.comCreateTrigger // Com_create_trigger
}

// ComCreateUser returns number of CREATE USER https://mariadb.com/kb/en/create-user/ commands executed.
func (s *Status) ComCreateUser() int64 {
	return s.pStatus.comCreateUser // Com_create_user
}

// ComDeallocSql returns number of DEALLOCATE https://mariadb.com/kb/en/deallocate-drop-prepared-statement/ commands executed.
func (s *Status) ComDeallocSql() int64 {
	return s.pStatus.comDeallocSql // Com_dealloc_sql
}

// ComDeleteMulti returns number of multi-table DELETE https://mariadb.com/kb/en/delete/ commands executed.
func (s *Status) ComDeleteMulti() int64 {
	return s.pStatus.comDeleteMulti // Com_delete_multi
}

// ComDropDb returns number of DROP DATABASE https://mariadb.com/kb/en/drop-database/ commands executed.
func (s *Status) ComDropDb() int64 {
	return s.pStatus.comDropDb // Com_drop_db
}

// ComDropFunction returns number of DROP FUNCTION https://mariadb.com/kb/en/drop-function/ commands executed.
func (s *Status) ComDropFunction() int64 {
	return s.pStatus.comDropFunction // Com_drop_function
}

// ComDropProcedure returns number of DROP PROCEDURE https://mariadb.com/kb/en/drop-procedure/ commands executed.
func (s *Status) ComDropProcedure() int64 {
	return s.pStatus.comDropProcedure // Com_drop_procedure
}

// ComDropServer returns number of DROP SERVER https://mariadb.com/kb/en/drop-server/ commands executed.
func (s *Status) ComDropServer() int64 {
	return s.pStatus.comDropServer // Com_drop_server
}

// ComDropTrigger returns number of DROP TRIGGER https://mariadb.com/kb/en/drop-trigger/ commands executed.
func (s *Status) ComDropTrigger() int64 {
	return s.pStatus.comDropTrigger // Com_drop_trigger
}

// ComDropView returns number of DROP VIEW https://mariadb.com/kb/en/drop-view/ commands executed.
func (s *Status) ComDropView() int64 {
	return s.pStatus.comDropView // Com_drop_view
}

// ComExecuteSql returns number of EXECUTE https://mariadb.com/kb/en/execute-statement/ statements executed.
func (s *Status) ComExecuteSql() int64 {
	return s.pStatus.comExecuteSql // Com_execute_sql
}

// ComGetDiagnostics returns number of GET DIAGNOSTICS https://mariadb.com/kb/en/get-diagnostics/ commands executed.
func (s *Status) ComGetDiagnostics() int64 {
	return s.pStatus.comGetDiagnostics // Com_get_diagnostics
}

// ComGrant returns number of GRANT https://mariadb.com/kb/en/grant/ commands executed.
func (s *Status) ComGrant() int64 {
	return s.pStatus.comGrant // Com_grant
}

// ComHaClose returns number of HANDLER https://mariadb.com/kb/en/handler-commands/ table_name CLOSE commands executed.
func (s *Status) ComHaClose() int64 {
	return s.pStatus.comHaClose // Com_ha_close
}

// ComHaRead returns number of HANDLER https://mariadb.com/kb/en/handler-commands/ table_name READ commands executed.
func (s *Status) ComHaRead() int64 {
	return s.pStatus.comHaRead // Com_ha_read
}

// ComInsert returns number of INSERT https://mariadb.com/kb/en/insert/ commands executed.
func (s *Status) ComInsert() int64 {
	return s.pStatus.comInsert // Com_insert
}

// ComInstallPlugin returns number of INSTALL PLUGIN https://mariadb.com/kb/en/install-plugin/ commands executed.
func (s *Status) ComInstallPlugin() int64 {
	return s.pStatus.comInstallPlugin // Com_install_plugin
}

// ComLoad returns number of LOAD commands executed.
func (s *Status) ComLoad() int64 {
	return s.pStatus.comLoad // Com_load
}

// ComLoadMasterTable returns
func (s *Status) ComLoadMasterTable() int64 {
	return s.pStatus.comLoadMasterTable // Com_load_master_table
}

// ComLockTables returns number of [lock-tables|LOCK TABLES]] commands executed.
func (s *Status) ComLockTables() int64 {
	return s.pStatus.comLockTables // Com_lock_tables
}

// ComPreloadKeys returns
func (s *Status) ComPreloadKeys() int64 {
	return s.pStatus.comPreloadKeys // Com_preload_keys
}

// ComPurge returns number of PURGE https://mariadb.com/kb/en/sql-commands-purge-logs/ commands executed.
func (s *Status) ComPurge() int64 {
	return s.pStatus.comPurge // Com_purge
}

// ComReleaseSavepoint returns number of RELEASE SAVEPOINT https://mariadb.com/kb/en/savepoint/ commands executed.
func (s *Status) ComReleaseSavepoint() int64 {
	return s.pStatus.comReleaseSavepoint // Com_release_savepoint
}

// ComRenameUser returns number of RENAME USER https://mariadb.com/kb/en/rename-user/ commands executed.
func (s *Status) ComRenameUser() int64 {
	return s.pStatus.comRenameUser // Com_rename_user
}

// ComReplace returns number of REPLACE https://mariadb.com/kb/en/replace/ commands executed.
func (s *Status) ComReplace() int64 {
	return s.pStatus.comReplace // Com_replace
}

// ComReset returns number of RESET https://mariadb.com/kb/en/reset/ commands executed.
func (s *Status) ComReset() int64 {
	return s.pStatus.comReset // Com_reset
}

// ComRestoreTable returns
func (s *Status) ComRestoreTable() int64 {
	return s.pStatus.comRestoreTable // Com_restore_table
}

// ComRevoke returns number of REVOKE https://mariadb.com/kb/en/revoke/ commands executed.
func (s *Status) ComRevoke() int64 {
	return s.pStatus.comRevoke // Com_revoke
}

// ComRevokeGrant returns number of REVOKE https://mariadb.com/kb/en/revoke/#roles role commands executed.
func (s *Status) ComRevokeGrant() int64 {
	return s.pStatus.comRevokeGrant // Com_revoke_grant
}

// ComRollback returns number of ROLLBACK https://mariadb.com/kb/en/rollback/ commands executed. Differs from Handler_rollback <#handler_rollback>, which is the number of transaction rollback requests given to a storage engine.
func (s *Status) ComRollback() int64 {
	return s.pStatus.comRollback // Com_rollback
}

// ComSavepoint returns number of SAVEPOINT https://mariadb.com/kb/en/savepoint/ commands executed. Differs from Handler_savepoint <#handler_savepoint>, which is the number of transaction savepoint creation requests.
func (s *Status) ComSavepoint() int64 {
	return s.pStatus.comSavepoint // Com_savepoint
}

// ComSetOption returns number of SET OPTION https://mariadb.com/kb/en/set/ commands executed.
func (s *Status) ComSetOption() int64 {
	return s.pStatus.comSetOption // Com_set_option
}

// ComShowAuthors returns number of SHOW AUTHORS https://mariadb.com/kb/en/show-authors/ commands executed.
func (s *Status) ComShowAuthors() int64 {
	return s.pStatus.comShowAuthors // Com_show_authors
}

// ComShowBinlogs returns number of SHOW BINARY LOGS https://mariadb.com/kb/en/show-binary-logs/ commands executed.
func (s *Status) ComShowBinlogs() int64 {
	return s.pStatus.comShowBinlogs // Com_show_binlogs
}

// ComShowClientStatistics returns number of SHOW CLIENT STATISTICS https://mariadb.com/kb/en/show-client_statistics/ commands executed.
func (s *Status) ComShowClientStatistics() int64 {
	return s.pStatus.comShowClientStatistics // Com_show_client_statistics
}

// ComShowCollations returns number of SHOW COLLATION https://mariadb.com/kb/en/show-collation/ commands executed.
func (s *Status) ComShowCollations() int64 {
	return s.pStatus.comShowCollations // Com_show_collations
}

// ComShowContributors returns number of SHOW CONTRIBUTORS https://mariadb.com/kb/en/show-contributors/ commands executed.
func (s *Status) ComShowContributors() int64 {
	return s.pStatus.comShowContributors // Com_show_contributors
}

// ComShowCreateEvent returns number of SHOW CREATE EVENT https://mariadb.com/kb/en/show-create-event/ commands executed.
func (s *Status) ComShowCreateEvent() int64 {
	return s.pStatus.comShowCreateEvent // Com_show_create_event
}

// ComShowCreateProc returns number of SHOW CREATE PROCEDURE https://mariadb.com/kb/en/show-create-procedure/ commands executed.
func (s *Status) ComShowCreateProc() int64 {
	return s.pStatus.comShowCreateProc // Com_show_create_proc
}

// ComShowCreateTrigger returns number of SHOW CREATE TRIGGER https://mariadb.com/kb/en/show-create-table/ commands executed.
func (s *Status) ComShowCreateTrigger() int64 {
	return s.pStatus.comShowCreateTrigger // Com_show_create_trigger
}

// ComShowEngineLogs returns number of SHOW ENGINE LOGS https://mariadb.com/kb/en/show-engine/ commands executed.
func (s *Status) ComShowEngineLogs() int64 {
	return s.pStatus.comShowEngineLogs // Com_show_engine_logs
}

// ComShowEngineStatus returns number of SHOW ENGINE STATUS https://mariadb.com/kb/en/show-engine/ commands executed.
func (s *Status) ComShowEngineStatus() int64 {
	return s.pStatus.comShowEngineStatus // Com_show_engine_status
}

// ComShowErrors returns number of SHOW ERRORS https://mariadb.com/kb/en/show-errors/ commands executed.
func (s *Status) ComShowErrors() int64 {
	return s.pStatus.comShowErrors // Com_show_errors
}

// ComShowFields returns number of SHOW COLUMNS https://mariadb.com/kb/en/show-columns/ or SHOW FIELDS commands executed.
func (s *Status) ComShowFields() int64 {
	return s.pStatus.comShowFields // Com_show_fields
}

// ComShowGrants returns number of SHOW GRANTS https://mariadb.com/kb/en/show-grants/ commands executed.
func (s *Status) ComShowGrants() int64 {
	return s.pStatus.comShowGrants // Com_show_grants
}

// ComShowIndexStatistics returns number of SHOW INDEX_STATISTICS https://mariadb.com/kb/en/show-index_statistics/ commands executed.
func (s *Status) ComShowIndexStatistics() int64 {
	return s.pStatus.comShowIndexStatistics // Com_show_index_statistics
}

// ComShowOpenTables returns number of SHOW OPEN TABLES https://mariadb.com/kb/en/show-open-tables/ commands executed.
func (s *Status) ComShowOpenTables() int64 {
	return s.pStatus.comShowOpenTables // Com_show_open_tables
}

// ComShowPrivileges returns number of SHOW PRIVILEGES https://mariadb.com/kb/en/show-privileges/ commands executed.
func (s *Status) ComShowPrivileges() int64 {
	return s.pStatus.comShowPrivileges // Com_show_privileges
}

// ComShowProcesslist returns number of SHOW PROCESSLIST https://mariadb.com/kb/en/show-processlist/ commands executed.
func (s *Status) ComShowProcesslist() int64 {
	return s.pStatus.comShowProcesslist // Com_show_processlist
}

// ComShowProfiles returns number of SHOW PROFILES https://mariadb.com/kb/en/show-profiles/ commands executed.
func (s *Status) ComShowProfiles() int64 {
	return s.pStatus.comShowProfiles // Com_show_profiles
}

// ComShowStatus returns number of SHOW STATUS https://mariadb.com/kb/en/show-status/ commands executed.
func (s *Status) ComShowStatus() int64 {
	return s.pStatus.comShowStatus // Com_show_status
}

// ComShowTableStatistics returns number of SHOW TABLE STATISTICS https://mariadb.com/kb/en/show-table_statistics/ commands executed.
func (s *Status) ComShowTableStatistics() int64 {
	return s.pStatus.comShowTableStatistics // Com_show_table_statistics
}

// ComShowTables returns number of SHOW TABLES https://mariadb.com/kb/en/show-tables/ commands executed.
func (s *Status) ComShowTables() int64 {
	return s.pStatus.comShowTables // Com_show_tables
}

// ComShowUserStatistics returns number of SHOW USER STATISTICS https://mariadb.com/kb/en/show-user_statistics/ commands executed.
func (s *Status) ComShowUserStatistics() int64 {
	return s.pStatus.comShowUserStatistics // Com_show_user_statistics
}

// ComShowWarnings returns number of SHOW WARNINGS https://mariadb.com/kb/en/show-warnings/ commands executed.
func (s *Status) ComShowWarnings() int64 {
	return s.pStatus.comShowWarnings // Com_show_warnings
}

// ComStmtClose returns number of prepared statements https://mariadb.com/kb/en/prepared-statements/ closed (deallocated or dropped </kb/en/deallocate-drop-prepared-statement/>).
func (s *Status) ComStmtClose() int64 {
	return s.pStatus.comStmtClose // Com_stmt_close
}

// ComStmtFetch returns number of prepared statements https://mariadb.com/kb/en/prepared-statements/ fetched.
func (s *Status) ComStmtFetch() int64 {
	return s.pStatus.comStmtFetch // Com_stmt_fetch
}

// ComStmtReprepare returns number of prepared statements https://mariadb.com/kb/en/prepared-statements/ reprepared.
func (s *Status) ComStmtReprepare() int64 {
	return s.pStatus.comStmtReprepare // Com_stmt_reprepare
}

// ComStmtSendLongData returns number of prepared statements https://mariadb.com/kb/en/prepared-statements/ where the parameter data has been sent in chunks (long data).
func (s *Status) ComStmtSendLongData() int64 {
	return s.pStatus.comStmtSendLongData // Com_stmt_send_long_data
}

// ComUninstallPlugin returns number of UNINSTALL PLUGIN https://mariadb.com/kb/en/uninstall-plugin/ commands executed.
func (s *Status) ComUninstallPlugin() int64 {
	return s.pStatus.comUninstallPlugin // Com_uninstall_plugin
}

// ComUpdate returns number of UPDATE https://mariadb.com/kb/en/update/ commands executed.
func (s *Status) ComUpdate() int64 {
	return s.pStatus.comUpdate // Com_update
}

// ComXaCommit returns number of XA statements committed.
func (s *Status) ComXaCommit() int64 {
	return s.pStatus.comXaCommit // Com_xa_commit
}

// ComXaPrepare returns number of XA statements prepared.
func (s *Status) ComXaPrepare() int64 {
	return s.pStatus.comXaPrepare // Com_xa_prepare
}

// ComXaRollback returns number of XA statements rolled back.
func (s *Status) ComXaRollback() int64 {
	return s.pStatus.comXaRollback // Com_xa_rollback
}

// Compression returns whether client-server traffic is compressed.
func (s *Status) Compression() int64 {
	return s.pStatus.compression // Compression
}

// ConnectionErrorsInternal returns number of refused connections due to internal server errors, for example out of memory errors, or failed thread starts.
func (s *Status) ConnectionErrorsInternal() int64 {
	return s.pStatus.connectionErrorsInternal // Connection_errors_internal
}

// ConnectionErrorsMaxConnections returns number of refused connections due to the max_connections https://mariadb.com/kb/en/server-system-variables/#max_connections limit being reached.
func (s *Status) ConnectionErrorsMaxConnections() int64 {
	return s.pStatus.connectionErrorsMaxConnections // Connection_errors_max_connections
}

// ConnectionErrorsPeerAddress returns number of errors while searching for the connecting client IP address.
func (s *Status) ConnectionErrorsPeerAddress() int64 {
	return s.pStatus.connectionErrorsPeerAddress // Connection_errors_peer_address
}

// ConnectionErrorsSelect returns number of errors during calls to select() or poll() on the listening port. The client would not necessarily have been rejected in these cases.
func (s *Status) ConnectionErrorsSelect() int64 {
	return s.pStatus.connectionErrorsSelect // Connection_errors_select
}

// ConnectionErrorsTcpwrap returns number of connections the libwrap library refused.
func (s *Status) ConnectionErrorsTcpwrap() int64 {
	return s.pStatus.connectionErrorsTcpwrap // Connection_errors_tcpwrap
}

// Connections returns number of connection attempts (both successful and unsuccessful)
func (s *Status) Connections() int64 {
	return s.pStatus.connections // Connections
}

// CreatedTmpDiskTables returns number of on-disk temporary tables created.
func (s *Status) CreatedTmpDiskTables() int64 {
	return s.pStatus.createdTmpDiskTables // Created_tmp_disk_tables
}

// CreatedTmpTables returns number of temporary tables created.
func (s *Status) CreatedTmpTables() int64 {
	return s.pStatus.createdTmpTables // Created_tmp_tables
}

// DelayedInsertThreads returns number of INSERT DELAYED https://mariadb.com/kb/en/insert-delayed/ threads.
func (s *Status) DelayedInsertThreads() int64 {
	return s.pStatus.delayedInsertThreads // Delayed_insert_threads
}

// EmptyQueries returns number of queries returning no results. Note this is not the same as Com_empty_query https://mariadb.com#com_empty_query.
func (s *Status) EmptyQueries() int64 {
	return s.pStatus.emptyQueries // Empty_queries
}

// ExecutedEvents returns number of times events created with CREATE EVENT https://mariadb.com/kb/en/create-event/ have executed. This differs from Com_create_event <#com_create_event> in that it is only incremented when the event has run, not when it executes.
func (s *Status) ExecutedEvents() int64 {
	return s.pStatus.executedEvents // Executed_events
}

// ExecutedTriggers returns number of times triggers created with CREATE TRIGGER https://mariadb.com/kb/en/create-trigger/ have executed. This differs from Com_create_trigger <#com_create_trigger> in that it is only incremented when the trigger has run, not when it executes.
func (s *Status) ExecutedTriggers() int64 {
	return s.pStatus.executedTriggers // Executed_triggers
}

// FeatureDynamicColumns returns number of times the COLUMN_CREATE() https://mariadb.com/kb/en/column_create/ function was used.
func (s *Status) FeatureDynamicColumns() int64 {
	return s.pStatus.featureDynamicColumns // Feature_dynamic_columns
}

// FeatureFulltext returns number of times the MATCH … AGAINST() https://mariadb.com/kb/en/match-against/ function was used.
func (s *Status) FeatureFulltext() int64 {
	return s.pStatus.featureFulltext // Feature_fulltext
}

// FeatureGis returns number of times a table with a any of the geometry https://mariadb.com/kb/en/geometry-types/ columns was opened.
func (s *Status) FeatureGis() int64 {
	return s.pStatus.featureGis // Feature_gis
}

// FeatureLocale returns number of times the @@lc_messages https://mariadb.com/kb/en/server-system-variables/#lc_messages variable was assigned into.
func (s *Status) FeatureLocale() int64 {
	return s.pStatus.featureLocale // Feature_locale
}

// FeatureSubquery returns number of subqueries (excluding subqueries in the FROM clause) used.
func (s *Status) FeatureSubquery() int64 {
	return s.pStatus.featureSubquery // Feature_subquery
}

// FeatureTimezone returns number of times an explicit timezone (excluding UTC https://mariadb.com/kb/en/coordinated-universal-time/ and SYSTEM) was specified.
func (s *Status) FeatureTimezone() int64 {
	return s.pStatus.featureTimezone // Feature_timezone
}

// FeatureTrigger returns number of triggers loaded.
func (s *Status) FeatureTrigger() int64 {
	return s.pStatus.featureTrigger // Feature_trigger
}

// FeatureXml returns number of times XML functions (EXTRACTVALUE() https://mariadb.com/kb/en/extractvalue/ and UPDATEXML() </kb/en/updatexml/>) were used.
func (s *Status) FeatureXml() int64 {
	return s.pStatus.featureXml // Feature_xml
}

// FlushCommands returns number of FLUSH https://mariadb.com/kb/en/flush/ statements executed, as well as due to internal server flush requests. This differs from Com_flush <#com_flush>, which simply counts FLUSH statements, not internal server flush operations.
func (s *Status) FlushCommands() int64 {
	return s.pStatus.flushCommands // Flush_commands
}

// HandlerDelete returns number of times rows have been deleted from tables. Differs from Com_delete https://mariadb.com#com_delete, which counts DELETE </kb/en/delete/> statements.
func (s *Status) HandlerDelete() int64 {
	return s.pStatus.handlerDelete // Handler_delete
}

// HandlerExternalLock returns incremented for each call to the external_lock() function, which generally occurs at the beginning and end of access to a table instance.
func (s *Status) HandlerExternalLock() int64 {
	return s.pStatus.handlerExternalLock // Handler_external_lock
}

// HandlerIcpAttempts returns number of times pushed index condition was checked. The smaller the ratio of Handler_icp_attempts to Handler_icp_match https://mariadb.com#handler_icp_match the better the filtering. See Index Condition Pushdown </kb/en/index-condition-pushdown/>.
func (s *Status) HandlerIcpAttempts() int64 {
	return s.pStatus.handlerIcpAttempts // Handler_icp_attempts
}

// HandlerIcpMatch returns number of times pushed index condition was matched. The smaller the ratio of Handler_icp_attempts https://mariadb.com#handler_icp_attempts to Handler_icp_match the better the filtering. See See Index Condition Pushdown </kb/en/index-condition-pushdown/>.
func (s *Status) HandlerIcpMatch() int64 {
	return s.pStatus.handlerIcpMatch // Handler_icp_match
}

// HandlerMrrInit returns counts how many MRR (multi-range read) scans were performed. See Multi Range Read optimization https://mariadb.com/kb/en/multi-range-read-optimization/.
func (s *Status) HandlerMrrInit() int64 {
	return s.pStatus.handlerMrrInit // Handler_mrr_init
}

// HandlerMrrKeyRefills returns number of times key buffer was refilled (not counting the initial fill). A non-zero value indicates there wasn't enough memory to do key sort-and-sweep passes in one go. See Multi Range Read optimization https://mariadb.com/kb/en/multi-range-read-optimization/.
func (s *Status) HandlerMrrKeyRefills() int64 {
	return s.pStatus.handlerMrrKeyRefills // Handler_mrr_key_refills
}

// HandlerMrrRowidRefills returns number of times rowid buffer was refilled (not counting the initial fill). A non-zero value indicates there wasn't enough memory to do rowid sort-and-sweep passes in one go. See Multi Range Read optimization https://mariadb.com/kb/en/multi-range-read-optimization/.
func (s *Status) HandlerMrrRowidRefills() int64 {
	return s.pStatus.handlerMrrRowidRefills // Handler_mrr_rowid_refills
}

// HandlerPrepare returns number of two-phase commit prepares.
func (s *Status) HandlerPrepare() int64 {
	return s.pStatus.handlerPrepare // Handler_prepare
}

// HandlerReadKey returns number of row read requests based on an index value. A high value indicates indexes are regularly being used, which is usually positive.
func (s *Status) HandlerReadKey() int64 {
	return s.pStatus.handlerReadKey // Handler_read_key
}

// HandlerReadNext returns number of requests to read the next row from an index (in order). Increments when doing an index scan or querying an index column with a range constraint.
func (s *Status) HandlerReadNext() int64 {
	return s.pStatus.handlerReadNext // Handler_read_next
}

// HandlerReadRnd returns number of requests to read a row based on its position. If this value is high, you may not be using joins that don't use indexes properly, or be doing many full table scans.
func (s *Status) HandlerReadRnd() int64 {
	return s.pStatus.handlerReadRnd // Handler_read_rnd
}

// HandlerReadRndNext returns number of requests to read the next row. A large number of these may indicate many table scans and improperly used indexes.
func (s *Status) HandlerReadRndNext() int64 {
	return s.pStatus.handlerReadRndNext // Handler_read_rnd_next
}

// HandlerSavepoint returns number of transaction savepoint creation requests. Differs from Com_savepoint https://mariadb.com#com_savepoint which is the number of SAVEPOINT </kb/en/savepoint/> commands executed.
func (s *Status) HandlerSavepoint() int64 {
	return s.pStatus.handlerSavepoint // Handler_savepoint
}

// HandlerTmpUpdate returns number of requests to update a row to a temporary table.
func (s *Status) HandlerTmpUpdate() int64 {
	return s.pStatus.handlerTmpUpdate // Handler_tmp_update
}

// HandlerTmpWrite returns number of requests to write a row to a temporary table.
func (s *Status) HandlerTmpWrite() int64 {
	return s.pStatus.handlerTmpWrite // Handler_tmp_write
}

// HandlerUpdate returns number of requests to update a row in a table. Since MariaDB 5.3 https://mariadb.com/kb/en/what-is-mariadb-53/, this no longer counts temporary tables - see Handler_tmp_update <#handler_tmp_update>.
func (s *Status) HandlerUpdate() int64 {
	return s.pStatus.handlerUpdate // Handler_update
}

// KeyBlocksNotFlushed returns number of key cache blocks which have been modified but not flushed to disk.
func (s *Status) KeyBlocksNotFlushed() int64 {
	return s.pStatus.keyBlocksNotFlushed // Key_blocks_not_flushed
}

// KeyBlocksUsed returns max number of key cache blocks which have been used simultaneously.
func (s *Status) KeyBlocksUsed() int64 {
	return s.pStatus.keyBlocksUsed // Key_blocks_used
}

// KeyReadRequests returns number of key cache block read requests. See Optimizing key_buffer_size https://mariadb.com/kb/en/optimizing-key_buffer_size/.
func (s *Status) KeyReadRequests() int64 {
	return s.pStatus.keyReadRequests // Key_read_requests
}

// KeyWriteRequests returns number of requests to write a block to the key cache.
func (s *Status) KeyWriteRequests() int64 {
	return s.pStatus.keyWriteRequests // Key_write_requests
}

// LastQueryCost returns the most recent query optimizer query cost calculation. Can not be calculated for complex queries, such as subqueries or UNION. It will be set to 0 for complex queries.
func (s *Status) LastQueryCost() int64 {
	return s.pStatus.lastQueryCost // Last_query_cost
}

// MaxUsedConnections returns max number of connections ever open at the same time.
func (s *Status) MaxUsedConnections() int64 {
	return s.pStatus.maxUsedConnections // Max_used_connections
}

// NotFlushedDelayedRows returns number of INSERT DELAYED https://mariadb.com/kb/en/insert-delayed/ rows waiting to be written.
func (s *Status) NotFlushedDelayedRows() int64 {
	return s.pStatus.notFlushedDelayedRows // Not_flushed_delayed_rows
}

// OpenStreams returns number of currently opened streams, usually log files.
func (s *Status) OpenStreams() int64 {
	return s.pStatus.openStreams // Open_streams
}

// OpenTables returns number of currently opened tables, excluding temporary tables.
func (s *Status) OpenTables() int64 {
	return s.pStatus.openTables // Open_tables
}

// OpenedTableDefinitions returns number of .frm files that have been cached.
func (s *Status) OpenedTableDefinitions() int64 {
	return s.pStatus.openedTableDefinitions // Opened_table_definitions
}

// OpenedViews returns number of views the server has opened.
func (s *Status) OpenedViews() int64 {
	return s.pStatus.openedViews // Opened_views
}

// OqgraphBoostVersion returns oQGRAPH https://mariadb.com/kb/en/oqgraph/ boost version.
func (s *Status) OqgraphBoostVersion() int64 {
	return s.pStatus.oqgraphBoostVersion // Oqgraph_boost_version
}

// PerformanceSchemaAccountsLost returns number of times a row could not be added to the performance schema accounts table due to it being full.
func (s *Status) PerformanceSchemaAccountsLost() int64 {
	return s.pStatus.performanceSchemaAccountsLost // Performance_schema_accounts_lost
}

// PerformanceSchemaCondClassesLost returns number of condition instruments that could not be loaded.
func (s *Status) PerformanceSchemaCondClassesLost() int64 {
	return s.pStatus.performanceSchemaCondClassesLost // Performance_schema_cond_classes_lost
}

// PerformanceSchemaCondInstancesLost returns number of instances a condition object could not be created.
func (s *Status) PerformanceSchemaCondInstancesLost() int64 {
	return s.pStatus.performanceSchemaCondInstancesLost // Performance_schema_cond_instances_lost
}

// PerformanceSchemaDigestLost returns
func (s *Status) PerformanceSchemaDigestLost() int64 {
	return s.pStatus.performanceSchemaDigestLost // Performance_schema_digest_lost
}

// PerformanceSchemaFileClassesLost returns number of file instruments that could not be loaded.
func (s *Status) PerformanceSchemaFileClassesLost() int64 {
	return s.pStatus.performanceSchemaFileClassesLost // Performance_schema_file_classes_lost
}

// PerformanceSchemaFileHandlesLost returns number of instances a file object could not be opened.
func (s *Status) PerformanceSchemaFileHandlesLost() int64 {
	return s.pStatus.performanceSchemaFileHandlesLost // Performance_schema_file_handles_lost
}

// PerformanceSchemaFileInstancesLost returns number of instances a file object could not be created.
func (s *Status) PerformanceSchemaFileInstancesLost() int64 {
	return s.pStatus.performanceSchemaFileInstancesLost // Performance_schema_file_instances_lost
}

// PerformanceSchemaHostsLost returns number of times a row could not be added to the performance schema hosts table due to it being full.
func (s *Status) PerformanceSchemaHostsLost() int64 {
	return s.pStatus.performanceSchemaHostsLost // Performance_schema_hosts_lost
}

// PerformanceSchemaLockerLost returns number of events not recorded, due to either being recursive, or having a deeper nested events stack than the implementation limit.
func (s *Status) PerformanceSchemaLockerLost() int64 {
	return s.pStatus.performanceSchemaLockerLost // Performance_schema_locker_lost
}

// PerformanceSchemaMutexClassesLost returns number of mutual exclusion instruments that could not be loaded.
func (s *Status) PerformanceSchemaMutexClassesLost() int64 {
	return s.pStatus.performanceSchemaMutexClassesLost // Performance_schema_mutex_classes_lost
}

// PerformanceSchemaMutexInstancesLost returns number of instances a mutual exclusion object could not be created.
func (s *Status) PerformanceSchemaMutexInstancesLost() int64 {
	return s.pStatus.performanceSchemaMutexInstancesLost // Performance_schema_mutex_instances_lost
}

// PerformanceSchemaRwlockClassesLost returns number of read/write lock instruments that could not be loaded.
func (s *Status) PerformanceSchemaRwlockClassesLost() int64 {
	return s.pStatus.performanceSchemaRwlockClassesLost // Performance_schema_rwlock_classes_lost
}

// PerformanceSchemaRwlockInstancesLost returns number of instances a read/write lock object could not be created.
func (s *Status) PerformanceSchemaRwlockInstancesLost() int64 {
	return s.pStatus.performanceSchemaRwlockInstancesLost // Performance_schema_rwlock_instances_lost
}

// PerformanceSchemaSocketClassesLost returns
func (s *Status) PerformanceSchemaSocketClassesLost() int64 {
	return s.pStatus.performanceSchemaSocketClassesLost // Performance_schema_socket_classes_lost
}

// PerformanceSchemaSocketInstancesLost returns number of instances a socket object could not be created.
func (s *Status) PerformanceSchemaSocketInstancesLost() int64 {
	return s.pStatus.performanceSchemaSocketInstancesLost // Performance_schema_socket_instances_lost
}

// PerformanceSchemaStageClassesLost returns number of stage event instruments that could not be loaded.
func (s *Status) PerformanceSchemaStageClassesLost() int64 {
	return s.pStatus.performanceSchemaStageClassesLost // Performance_schema_stage_classes_lost
}

// PerformanceSchemaStatementClassesLost returns number of statement instruments that could not be loaded.
func (s *Status) PerformanceSchemaStatementClassesLost() int64 {
	return s.pStatus.performanceSchemaStatementClassesLost // Performance_schema_statement_classes_lost
}

// PerformanceSchemaTableHandlesLost returns number of instances a table object could not be opened.
func (s *Status) PerformanceSchemaTableHandlesLost() int64 {
	return s.pStatus.performanceSchemaTableHandlesLost // Performance_schema_table_handles_lost
}

// PerformanceSchemaTableInstancesLost returns number of instances a table object could not be created.
func (s *Status) PerformanceSchemaTableInstancesLost() int64 {
	return s.pStatus.performanceSchemaTableInstancesLost // Performance_schema_table_instances_lost
}

// PerformanceSchemaThreadClassesLost returns number of thread instruments that could not be loaded.
func (s *Status) PerformanceSchemaThreadClassesLost() int64 {
	return s.pStatus.performanceSchemaThreadClassesLost // Performance_schema_thread_classes_lost
}

// PerformanceSchemaThreadInstancesLost returns number of instances thread object could not be created.
func (s *Status) PerformanceSchemaThreadInstancesLost() int64 {
	return s.pStatus.performanceSchemaThreadInstancesLost // Performance_schema_thread_instances_lost
}

// PerformanceSchemaUsersLost returns number of times a row could not be added to the performance schema users table due to it being full.
func (s *Status) PerformanceSchemaUsersLost() int64 {
	return s.pStatus.performanceSchemaUsersLost // Performance_schema_users_lost
}

// PreparedStmtCount returns current number of prepared statements.
func (s *Status) PreparedStmtCount() int64 {
	return s.pStatus.preparedStmtCount // Prepared_stmt_count
}

// QcacheFreeBlocks returns number of free query cache https://mariadb.com/kb/en/query-cache/ memory blocks.
func (s *Status) QcacheFreeBlocks() int64 {
	return s.pStatus.qcacheFreeBlocks // Qcache_free_blocks
}

// QcacheHits returns number of requests served by the query cache https://mariadb.com/kb/en/query-cache/.
func (s *Status) QcacheHits() int64 {
	return s.pStatus.qcacheHits // Qcache_hits
}

// QcacheLowmemPrunes returns number of pruning operations performed to remove old results to make space for new results in the query cache https://mariadb.com/kb/en/query-cache/.
func (s *Status) QcacheLowmemPrunes() int64 {
	return s.pStatus.qcacheLowmemPrunes // Qcache_lowmem_prunes
}

// QcacheQueriesInCache returns number of queries currently cached by the query cache https://mariadb.com/kb/en/query-cache/.
func (s *Status) QcacheQueriesInCache() int64 {
	return s.pStatus.qcacheQueriesInCache // Qcache_queries_in_cache
}

// Queries returns number of statements executed by the server, excluding COM_PING and COM_STATISTICS. Differs from Questions https://mariadb.com#questions in that it also counts statements executed within stored programs </kb/en/stored-programs-and-views/>.
func (s *Status) Queries() int64 {
	return s.pStatus.queries // Queries
}

// RowsRead returns number of requests to read a row (excluding temporary tables).
func (s *Status) RowsRead() int64 {
	return s.pStatus.rowsRead // Rows_read
}

// RowsSent returns
func (s *Status) RowsSent() int64 {
	return s.pStatus.rowsSent // Rows_sent
}

// RowsTmpRead returns number of requests to read a row in a temporary table.
func (s *Status) RowsTmpRead() int64 {
	return s.pStatus.rowsTmpRead // Rows_tmp_read
}

// RplRecoveryRank returns unused
func (s *Status) RplRecoveryRank() int64 {
	return s.pStatus.rplRecoveryRank // Rpl_recovery_rank
}

// SelectFullRangeJoin returns number of joins which used a range search of the first table.
func (s *Status) SelectFullRangeJoin() int64 {
	return s.pStatus.selectFullRangeJoin // Select_full_range_join
}

// SelectRangeCheck returns number of joins without keys that check for key usage after each row. If not zero, you may need to check table indexes.
func (s *Status) SelectRangeCheck() int64 {
	return s.pStatus.selectRangeCheck // Select_range_check
}

// SlowLaunchThreads returns number of threads which took longer than slow_launch_time https://mariadb.com/kb/en/server-system-variables/#slow_launch_time to create.
func (s *Status) SlowLaunchThreads() int64 {
	return s.pStatus.slowLaunchThreads // Slow_launch_threads
}

// SortMergePasses returns number of merge passes performed by the sort algorithm. If too high, you may need to increase the sort_buffer_size https://mariadb.com/kb/en/server-system-variables/#sort_buffer_size.
func (s *Status) SortMergePasses() int64 {
	return s.pStatus.sortMergePasses // Sort_merge_passes
}

// SortRows returns number of rows sorted.
func (s *Status) SortRows() int64 {
	return s.pStatus.sortRows // Sort_rows
}

// SphinxError returns see SHOW ENGINE SPHINX STATUS https://mariadb.com/kb/en/about-sphinxse/#show-engine-sphinx-status.
func (s *Status) SphinxError() int64 {
	return s.pStatus.sphinxError // Sphinx_error
}

// SphinxTime returns see SHOW ENGINE SPHINX STATUS https://mariadb.com/kb/en/about-sphinxse/#show-engine-sphinx-status.
func (s *Status) SphinxTime() int64 {
	return s.pStatus.sphinxTime // Sphinx_time
}

// SphinxTotal returns see SHOW ENGINE SPHINX STATUS https://mariadb.com/kb/en/about-sphinxse/#show-engine-sphinx-status.
func (s *Status) SphinxTotal() int64 {
	return s.pStatus.sphinxTotal // Sphinx_total
}

// SphinxTotalFound returns see SHOW ENGINE SPHINX STATUS https://mariadb.com/kb/en/about-sphinxse/#show-engine-sphinx-status.
func (s *Status) SphinxTotalFound() int64 {
	return s.pStatus.sphinxTotalFound // Sphinx_total_found
}

// SphinxWordCount returns see SHOW ENGINE SPHINX STATUS https://mariadb.com/kb/en/about-sphinxse/#show-engine-sphinx-status.
func (s *Status) SphinxWordCount() int64 {
	return s.pStatus.sphinxWordCount // Sphinx_word_count
}

// SphinxWords returns see SHOW ENGINE SPHINX STATUS https://mariadb.com/kb/en/about-sphinxse/#show-engine-sphinx-status.
func (s *Status) SphinxWords() int64 {
	return s.pStatus.sphinxWords // Sphinx_words
}

// SslAcceptRenegotiates returns number of negotiates needed to establish the SSL connection.
func (s *Status) SslAcceptRenegotiates() int64 {
	return s.pStatus.sslAcceptRenegotiates // Ssl_accept_renegotiates
}

// SslCallbackCacheHits returns number of sessions retrieved from session cache.
func (s *Status) SslCallbackCacheHits() int64 {
	return s.pStatus.sslCallbackCacheHits // Ssl_callback_cache_hits
}

// SslCipherList returns list of available SSL ciphers.
func (s *Status) SslCipherList() int64 {
	return s.pStatus.sslCipherList // Ssl_cipher_list
}

// SslConnectRenegotiates returns number of negotiates needed to establish the connection to an SSL-enabled master.
func (s *Status) SslConnectRenegotiates() int64 {
	return s.pStatus.sslConnectRenegotiates // Ssl_connect_renegotiates
}

// SslCtxVerifyMode returns sSL context verification mode.
func (s *Status) SslCtxVerifyMode() int64 {
	return s.pStatus.sslCtxVerifyMode // Ssl_ctx_verify_mode
}

// SslFinishedAccepts returns number of successful SSL/TLS sessions in server mode.
func (s *Status) SslFinishedAccepts() int64 {
	return s.pStatus.sslFinishedAccepts // Ssl_finished_accepts
}

// SslServerNotAfter returns last valid date for the SSL certificate.
func (s *Status) SslServerNotAfter() int64 {
	return s.pStatus.sslServerNotAfter // Ssl_server_not_after
}

// SslSessionCacheHits returns number of SSL sessions found in the session cache.
func (s *Status) SslSessionCacheHits() int64 {
	return s.pStatus.sslSessionCacheHits // Ssl_session_cache_hits
}

// SslSessionCacheMode returns mode used for SSL caching by the server.
func (s *Status) SslSessionCacheMode() int64 {
	return s.pStatus.sslSessionCacheMode // Ssl_session_cache_mode
}

// SslSessionCacheSize returns size of the session cache.
func (s *Status) SslSessionCacheSize() int64 {
	return s.pStatus.sslSessionCacheSize // Ssl_session_cache_size
}

// SslSessionsReused returns number of sessions reused.
func (s *Status) SslSessionsReused() int64 {
	return s.pStatus.sslSessionsReused // Ssl_sessions_reused
}

// SslVerifyDepth returns sSL verification depth.
func (s *Status) SslVerifyDepth() int64 {
	return s.pStatus.sslVerifyDepth // Ssl_verify_depth
}

// SslVersion returns sSL version in use.
func (s *Status) SslVersion() int64 {
	return s.pStatus.sslVersion // Ssl_version
}

// SubqueryCacheMiss returns counter for all subquery cache https://mariadb.com/kb/en/subquery-cache/ misses
func (s *Status) SubqueryCacheMiss() int64 {
	return s.pStatus.subqueryCacheMiss // Subquery_cache_miss
}

// Syncs returns number of times my_sync() has been called, or the number of times the server has had to force data to disk. Covers the binary log https://mariadb.com/kb/en/binary-log/, .frm creation (if these operations are configured to sync) and some storage engines (Archive </kb/en/archive/>, CSV </kb/en/csv/>, Aria </kb/en/aria/>), but not XtraDB/InnoDB </kb/en/innodb/>).
func (s *Status) Syncs() int64 {
	return s.pStatus.syncs // Syncs
}

// TableLocksImmediate returns number of table locks which were completed immediately.
func (s *Status) TableLocksImmediate() int64 {
	return s.pStatus.tableLocksImmediate // Table_locks_immediate
}

// TcLogMaxPagesUsed returns max number of pages used by the transaction coordinator recovery log.
func (s *Status) TcLogMaxPagesUsed() int64 {
	return s.pStatus.tcLogMaxPagesUsed // Tc_log_max_pages_used
}

// TcLogPageWaits returns number of times a two-phase commit was forced to wait for a free transaction coordinator recovery log page.
func (s *Status) TcLogPageWaits() int64 {
	return s.pStatus.tcLogPageWaits // Tc_log_page_waits
}

// ThreadpoolThreads returns number of threads in the thread pool https://mariadb.com/kb/en/threadpool-in-55/.
func (s *Status) ThreadpoolThreads() int64 {
	return s.pStatus.threadpoolThreads // Threadpool_threads
}

// ThreadsCached returns number of threads cached in the thread cache.
func (s *Status) ThreadsCached() int64 {
	return s.pStatus.threadsCached // Threads_cached
}

// ThreadsCreated returns number of threads created to respond to client connections. If too large, look at increasing thread_cache_size https://mariadb.com/kb/en/server-system-variables/#thread_cache_size.
func (s *Status) ThreadsCreated() int64 {
	return s.pStatus.threadsCreated // Threads_created
}

// Uptime returns number of seconds the server has been running.
func (s *Status) Uptime() int64 {
	return s.pStatus.uptime // Uptime
}

var ErrInvalidCast = errors.New("invalid cast")

func setInt64(src interface{}, dst *int64) error {
	val, ok := src.(int64)
	if !ok {
		return ErrInvalidCast
	}
	*dst = val
	return nil
}

func ReadStatus(db *sql.DB) (status Status, err error) {
	rows, err := db.Query("SHOW /*!50002 GLOBAL */ STATUS")
	if err != nil {
		return status, err
	}

	for rows.Next() {
		var name string
		var val interface{}
		err = rows.Scan(&name, &val)
		if err != nil {
			return status, err
		}
		switch name {
		case "Aborted_connects":
			err = setInt64(val, &status.pStatus.AbortedConnects)
		case "Binlog_bytes_written":
			err = setInt64(val, &status.pStatus.BinlogBytesWritten)
		case "Binlog_cache_disk_use":
			err = setInt64(val, &status.pStatus.BinlogCacheDiskUse)
		case "Binlog_commits":
			err = setInt64(val, &status.pStatus.BinlogCommits)
		case "Binlog_group_commits":
			err = setInt64(val, &status.pStatus.BinlogGroupCommits)
		case "Binlog_snapshot_files":
			err = setInt64(val, &status.pStatus.BinlogSnapshotFiles)
		case "Binlog_snapshot_position":
			err = setInt64(val, &status.pStatus.BinlogSnapshotPosition)
		case "Binlog_stmt_cache_disk_use":
			err = setInt64(val, &status.pStatus.BinlogStmtCacheDiskUse)
		case "Binlog_stmt_cache_use":
			err = setInt64(val, &status.pStatus.BinlogStmtCacheUse)
		case "Busy_time":
			err = setInt64(val, &status.pStatus.BusyTime)
		case "Bytes_received":
			err = setInt64(val, &status.pStatus.BytesReceived)
		case "Com_admin_commands":
			err = setInt64(val, &status.pStatus.ComAdminCommands)
		case "Com_alter_db_upgrade":
			err = setInt64(val, &status.pStatus.ComAlterDbUpgrade)
		case "Com_alter_function":
			err = setInt64(val, &status.pStatus.ComAlterFunction)
		case "Com_alter_server":
			err = setInt64(val, &status.pStatus.ComAlterServer)
		case "Com_alter_tablespace":
			err = setInt64(val, &status.pStatus.ComAlterTablespace)
		case "Com_assign_to_keycache":
			err = setInt64(val, &status.pStatus.ComAssignToKeycache)
		case "Com_begin":
			err = setInt64(val, &status.pStatus.ComBegin)
		case "Com_call_procedure":
			err = setInt64(val, &status.pStatus.ComCallProcedure)
		case "Com_check":
			err = setInt64(val, &status.pStatus.ComCheck)
		case "Com_commit":
			err = setInt64(val, &status.pStatus.ComCommit)
		case "Com_create_event":
			err = setInt64(val, &status.pStatus.ComCreateEvent)
		case "Com_create_index":
			err = setInt64(val, &status.pStatus.ComCreateIndex)
		case "Com_create_role":
			err = setInt64(val, &status.pStatus.ComCreateRole)
		case "Com_create_server":
			err = setInt64(val, &status.pStatus.ComCreateServer)
		case "Com_create_trigger":
			err = setInt64(val, &status.pStatus.ComCreateTrigger)
		case "Com_create_user":
			err = setInt64(val, &status.pStatus.ComCreateUser)
		case "Com_dealloc_sql":
			err = setInt64(val, &status.pStatus.ComDeallocSql)
		case "Com_delete_multi":
			err = setInt64(val, &status.pStatus.ComDeleteMulti)
		case "Com_drop_db":
			err = setInt64(val, &status.pStatus.ComDropDb)
		case "Com_drop_function":
			err = setInt64(val, &status.pStatus.ComDropFunction)
		case "Com_drop_procedure":
			err = setInt64(val, &status.pStatus.ComDropProcedure)
		case "Com_drop_server":
			err = setInt64(val, &status.pStatus.ComDropServer)
		case "Com_drop_trigger":
			err = setInt64(val, &status.pStatus.ComDropTrigger)
		case "Com_drop_view":
			err = setInt64(val, &status.pStatus.ComDropView)
		case "Com_execute_sql":
			err = setInt64(val, &status.pStatus.ComExecuteSql)
		case "Com_get_diagnostics":
			err = setInt64(val, &status.pStatus.ComGetDiagnostics)
		case "Com_grant":
			err = setInt64(val, &status.pStatus.ComGrant)
		case "Com_ha_close":
			err = setInt64(val, &status.pStatus.ComHaClose)
		case "Com_ha_read":
			err = setInt64(val, &status.pStatus.ComHaRead)
		case "Com_insert":
			err = setInt64(val, &status.pStatus.ComInsert)
		case "Com_install_plugin":
			err = setInt64(val, &status.pStatus.ComInstallPlugin)
		case "Com_load":
			err = setInt64(val, &status.pStatus.ComLoad)
		case "Com_load_master_table":
			err = setInt64(val, &status.pStatus.ComLoadMasterTable)
		case "Com_lock_tables":
			err = setInt64(val, &status.pStatus.ComLockTables)
		case "Com_preload_keys":
			err = setInt64(val, &status.pStatus.ComPreloadKeys)
		case "Com_purge":
			err = setInt64(val, &status.pStatus.ComPurge)
		case "Com_release_savepoint":
			err = setInt64(val, &status.pStatus.ComReleaseSavepoint)
		case "Com_rename_user":
			err = setInt64(val, &status.pStatus.ComRenameUser)
		case "Com_replace":
			err = setInt64(val, &status.pStatus.ComReplace)
		case "Com_reset":
			err = setInt64(val, &status.pStatus.ComReset)
		case "Com_restore_table":
			err = setInt64(val, &status.pStatus.ComRestoreTable)
		case "Com_revoke":
			err = setInt64(val, &status.pStatus.ComRevoke)
		case "Com_revoke_grant":
			err = setInt64(val, &status.pStatus.ComRevokeGrant)
		case "Com_rollback":
			err = setInt64(val, &status.pStatus.ComRollback)
		case "Com_savepoint":
			err = setInt64(val, &status.pStatus.ComSavepoint)
		case "Com_set_option":
			err = setInt64(val, &status.pStatus.ComSetOption)
		case "Com_show_authors":
			err = setInt64(val, &status.pStatus.ComShowAuthors)
		case "Com_show_binlogs":
			err = setInt64(val, &status.pStatus.ComShowBinlogs)
		case "Com_show_client_statistics":
			err = setInt64(val, &status.pStatus.ComShowClientStatistics)
		case "Com_show_collations":
			err = setInt64(val, &status.pStatus.ComShowCollations)
		case "Com_show_contributors":
			err = setInt64(val, &status.pStatus.ComShowContributors)
		case "Com_show_create_event":
			err = setInt64(val, &status.pStatus.ComShowCreateEvent)
		case "Com_show_create_proc":
			err = setInt64(val, &status.pStatus.ComShowCreateProc)
		case "Com_show_create_trigger":
			err = setInt64(val, &status.pStatus.ComShowCreateTrigger)
		case "Com_show_engine_logs":
			err = setInt64(val, &status.pStatus.ComShowEngineLogs)
		case "Com_show_engine_status":
			err = setInt64(val, &status.pStatus.ComShowEngineStatus)
		case "Com_show_errors":
			err = setInt64(val, &status.pStatus.ComShowErrors)
		case "Com_show_fields":
			err = setInt64(val, &status.pStatus.ComShowFields)
		case "Com_show_grants":
			err = setInt64(val, &status.pStatus.ComShowGrants)
		case "Com_show_index_statistics":
			err = setInt64(val, &status.pStatus.ComShowIndexStatistics)
		case "Com_show_open_tables":
			err = setInt64(val, &status.pStatus.ComShowOpenTables)
		case "Com_show_privileges":
			err = setInt64(val, &status.pStatus.ComShowPrivileges)
		case "Com_show_processlist":
			err = setInt64(val, &status.pStatus.ComShowProcesslist)
		case "Com_show_profiles":
			err = setInt64(val, &status.pStatus.ComShowProfiles)
		case "Com_show_status":
			err = setInt64(val, &status.pStatus.ComShowStatus)
		case "Com_show_table_statistics":
			err = setInt64(val, &status.pStatus.ComShowTableStatistics)
		case "Com_show_tables":
			err = setInt64(val, &status.pStatus.ComShowTables)
		case "Com_show_user_statistics":
			err = setInt64(val, &status.pStatus.ComShowUserStatistics)
		case "Com_show_warnings":
			err = setInt64(val, &status.pStatus.ComShowWarnings)
		case "Com_stmt_close":
			err = setInt64(val, &status.pStatus.ComStmtClose)
		case "Com_stmt_fetch":
			err = setInt64(val, &status.pStatus.ComStmtFetch)
		case "Com_stmt_reprepare":
			err = setInt64(val, &status.pStatus.ComStmtReprepare)
		case "Com_stmt_send_long_data":
			err = setInt64(val, &status.pStatus.ComStmtSendLongData)
		case "Com_uninstall_plugin":
			err = setInt64(val, &status.pStatus.ComUninstallPlugin)
		case "Com_update":
			err = setInt64(val, &status.pStatus.ComUpdate)
		case "Com_xa_commit":
			err = setInt64(val, &status.pStatus.ComXaCommit)
		case "Com_xa_prepare":
			err = setInt64(val, &status.pStatus.ComXaPrepare)
		case "Com_xa_rollback":
			err = setInt64(val, &status.pStatus.ComXaRollback)
		case "Compression":
			err = setInt64(val, &status.pStatus.Compression)
		case "Connection_errors_internal":
			err = setInt64(val, &status.pStatus.ConnectionErrorsInternal)
		case "Connection_errors_max_connections":
			err = setInt64(val, &status.pStatus.ConnectionErrorsMaxConnections)
		case "Connection_errors_peer_address":
			err = setInt64(val, &status.pStatus.ConnectionErrorsPeerAddress)
		case "Connection_errors_select":
			err = setInt64(val, &status.pStatus.ConnectionErrorsSelect)
		case "Connection_errors_tcpwrap":
			err = setInt64(val, &status.pStatus.ConnectionErrorsTcpwrap)
		case "Connections":
			err = setInt64(val, &status.pStatus.Connections)
		case "Created_tmp_disk_tables":
			err = setInt64(val, &status.pStatus.CreatedTmpDiskTables)
		case "Created_tmp_tables":
			err = setInt64(val, &status.pStatus.CreatedTmpTables)
		case "Delayed_insert_threads":
			err = setInt64(val, &status.pStatus.DelayedInsertThreads)
		case "Empty_queries":
			err = setInt64(val, &status.pStatus.EmptyQueries)
		case "Executed_events":
			err = setInt64(val, &status.pStatus.ExecutedEvents)
		case "Executed_triggers":
			err = setInt64(val, &status.pStatus.ExecutedTriggers)
		case "Feature_dynamic_columns":
			err = setInt64(val, &status.pStatus.FeatureDynamicColumns)
		case "Feature_fulltext":
			err = setInt64(val, &status.pStatus.FeatureFulltext)
		case "Feature_gis":
			err = setInt64(val, &status.pStatus.FeatureGis)
		case "Feature_locale":
			err = setInt64(val, &status.pStatus.FeatureLocale)
		case "Feature_subquery":
			err = setInt64(val, &status.pStatus.FeatureSubquery)
		case "Feature_timezone":
			err = setInt64(val, &status.pStatus.FeatureTimezone)
		case "Feature_trigger":
			err = setInt64(val, &status.pStatus.FeatureTrigger)
		case "Feature_xml":
			err = setInt64(val, &status.pStatus.FeatureXml)
		case "Flush_commands":
			err = setInt64(val, &status.pStatus.FlushCommands)
		case "Handler_delete":
			err = setInt64(val, &status.pStatus.HandlerDelete)
		case "Handler_external_lock":
			err = setInt64(val, &status.pStatus.HandlerExternalLock)
		case "Handler_icp_attempts":
			err = setInt64(val, &status.pStatus.HandlerIcpAttempts)
		case "Handler_icp_match":
			err = setInt64(val, &status.pStatus.HandlerIcpMatch)
		case "Handler_mrr_init":
			err = setInt64(val, &status.pStatus.HandlerMrrInit)
		case "Handler_mrr_key_refills":
			err = setInt64(val, &status.pStatus.HandlerMrrKeyRefills)
		case "Handler_mrr_rowid_refills":
			err = setInt64(val, &status.pStatus.HandlerMrrRowidRefills)
		case "Handler_prepare":
			err = setInt64(val, &status.pStatus.HandlerPrepare)
		case "Handler_read_key":
			err = setInt64(val, &status.pStatus.HandlerReadKey)
		case "Handler_read_next":
			err = setInt64(val, &status.pStatus.HandlerReadNext)
		case "Handler_read_rnd":
			err = setInt64(val, &status.pStatus.HandlerReadRnd)
		case "Handler_read_rnd_next":
			err = setInt64(val, &status.pStatus.HandlerReadRndNext)
		case "Handler_savepoint":
			err = setInt64(val, &status.pStatus.HandlerSavepoint)
		case "Handler_tmp_update":
			err = setInt64(val, &status.pStatus.HandlerTmpUpdate)
		case "Handler_tmp_write":
			err = setInt64(val, &status.pStatus.HandlerTmpWrite)
		case "Handler_update":
			err = setInt64(val, &status.pStatus.HandlerUpdate)
		case "Key_blocks_not_flushed":
			err = setInt64(val, &status.pStatus.KeyBlocksNotFlushed)
		case "Key_blocks_used":
			err = setInt64(val, &status.pStatus.KeyBlocksUsed)
		case "Key_read_requests":
			err = setInt64(val, &status.pStatus.KeyReadRequests)
		case "Key_write_requests":
			err = setInt64(val, &status.pStatus.KeyWriteRequests)
		case "Last_query_cost":
			err = setInt64(val, &status.pStatus.LastQueryCost)
		case "Max_used_connections":
			err = setInt64(val, &status.pStatus.MaxUsedConnections)
		case "Not_flushed_delayed_rows":
			err = setInt64(val, &status.pStatus.NotFlushedDelayedRows)
		case "Open_streams":
			err = setInt64(val, &status.pStatus.OpenStreams)
		case "Open_tables":
			err = setInt64(val, &status.pStatus.OpenTables)
		case "Opened_table_definitions":
			err = setInt64(val, &status.pStatus.OpenedTableDefinitions)
		case "Opened_views":
			err = setInt64(val, &status.pStatus.OpenedViews)
		case "Oqgraph_boost_version":
			err = setInt64(val, &status.pStatus.OqgraphBoostVersion)
		case "Performance_schema_accounts_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaAccountsLost)
		case "Performance_schema_cond_classes_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaCondClassesLost)
		case "Performance_schema_cond_instances_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaCondInstancesLost)
		case "Performance_schema_digest_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaDigestLost)
		case "Performance_schema_file_classes_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaFileClassesLost)
		case "Performance_schema_file_handles_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaFileHandlesLost)
		case "Performance_schema_file_instances_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaFileInstancesLost)
		case "Performance_schema_hosts_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaHostsLost)
		case "Performance_schema_locker_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaLockerLost)
		case "Performance_schema_mutex_classes_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaMutexClassesLost)
		case "Performance_schema_mutex_instances_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaMutexInstancesLost)
		case "Performance_schema_rwlock_classes_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaRwlockClassesLost)
		case "Performance_schema_rwlock_instances_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaRwlockInstancesLost)
		case "Performance_schema_socket_classes_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaSocketClassesLost)
		case "Performance_schema_socket_instances_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaSocketInstancesLost)
		case "Performance_schema_stage_classes_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaStageClassesLost)
		case "Performance_schema_statement_classes_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaStatementClassesLost)
		case "Performance_schema_table_handles_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaTableHandlesLost)
		case "Performance_schema_table_instances_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaTableInstancesLost)
		case "Performance_schema_thread_classes_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaThreadClassesLost)
		case "Performance_schema_thread_instances_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaThreadInstancesLost)
		case "Performance_schema_users_lost":
			err = setInt64(val, &status.pStatus.PerformanceSchemaUsersLost)
		case "Prepared_stmt_count":
			err = setInt64(val, &status.pStatus.PreparedStmtCount)
		case "Qcache_free_blocks":
			err = setInt64(val, &status.pStatus.QcacheFreeBlocks)
		case "Qcache_hits":
			err = setInt64(val, &status.pStatus.QcacheHits)
		case "Qcache_lowmem_prunes":
			err = setInt64(val, &status.pStatus.QcacheLowmemPrunes)
		case "Qcache_queries_in_cache":
			err = setInt64(val, &status.pStatus.QcacheQueriesInCache)
		case "Queries":
			err = setInt64(val, &status.pStatus.Queries)
		case "Rows_read":
			err = setInt64(val, &status.pStatus.RowsRead)
		case "Rows_sent":
			err = setInt64(val, &status.pStatus.RowsSent)
		case "Rows_tmp_read":
			err = setInt64(val, &status.pStatus.RowsTmpRead)
		case "Rpl_recovery_rank":
			err = setInt64(val, &status.pStatus.RplRecoveryRank)
		case "Select_full_range_join":
			err = setInt64(val, &status.pStatus.SelectFullRangeJoin)
		case "Select_range_check":
			err = setInt64(val, &status.pStatus.SelectRangeCheck)
		case "Slow_launch_threads":
			err = setInt64(val, &status.pStatus.SlowLaunchThreads)
		case "Sort_merge_passes":
			err = setInt64(val, &status.pStatus.SortMergePasses)
		case "Sort_rows":
			err = setInt64(val, &status.pStatus.SortRows)
		case "Sphinx_error":
			err = setInt64(val, &status.pStatus.SphinxError)
		case "Sphinx_time":
			err = setInt64(val, &status.pStatus.SphinxTime)
		case "Sphinx_total":
			err = setInt64(val, &status.pStatus.SphinxTotal)
		case "Sphinx_total_found":
			err = setInt64(val, &status.pStatus.SphinxTotalFound)
		case "Sphinx_word_count":
			err = setInt64(val, &status.pStatus.SphinxWordCount)
		case "Sphinx_words":
			err = setInt64(val, &status.pStatus.SphinxWords)
		case "Ssl_accept_renegotiates":
			err = setInt64(val, &status.pStatus.SslAcceptRenegotiates)
		case "Ssl_callback_cache_hits":
			err = setInt64(val, &status.pStatus.SslCallbackCacheHits)
		case "Ssl_cipher_list":
			err = setInt64(val, &status.pStatus.SslCipherList)
		case "Ssl_connect_renegotiates":
			err = setInt64(val, &status.pStatus.SslConnectRenegotiates)
		case "Ssl_ctx_verify_mode":
			err = setInt64(val, &status.pStatus.SslCtxVerifyMode)
		case "Ssl_finished_accepts":
			err = setInt64(val, &status.pStatus.SslFinishedAccepts)
		case "Ssl_server_not_after":
			err = setInt64(val, &status.pStatus.SslServerNotAfter)
		case "Ssl_session_cache_hits":
			err = setInt64(val, &status.pStatus.SslSessionCacheHits)
		case "Ssl_session_cache_mode":
			err = setInt64(val, &status.pStatus.SslSessionCacheMode)
		case "Ssl_session_cache_size":
			err = setInt64(val, &status.pStatus.SslSessionCacheSize)
		case "Ssl_sessions_reused":
			err = setInt64(val, &status.pStatus.SslSessionsReused)
		case "Ssl_verify_depth":
			err = setInt64(val, &status.pStatus.SslVerifyDepth)
		case "Ssl_version":
			err = setInt64(val, &status.pStatus.SslVersion)
		case "Subquery_cache_miss":
			err = setInt64(val, &status.pStatus.SubqueryCacheMiss)
		case "Syncs":
			err = setInt64(val, &status.pStatus.Syncs)
		case "Table_locks_immediate":
			err = setInt64(val, &status.pStatus.TableLocksImmediate)
		case "Tc_log_max_pages_used":
			err = setInt64(val, &status.pStatus.TcLogMaxPagesUsed)
		case "Tc_log_page_waits":
			err = setInt64(val, &status.pStatus.TcLogPageWaits)
		case "Threadpool_threads":
			err = setInt64(val, &status.pStatus.ThreadpoolThreads)
		case "Threads_cached":
			err = setInt64(val, &status.pStatus.ThreadsCached)
		case "Threads_created":
			err = setInt64(val, &status.pStatus.ThreadsCreated)
		case "Uptime":
			err = setInt64(val, &status.pStatus.Uptime)
		default:
		}
		if err != nil {
			return status, err
		}
	}
	return status, err
}
