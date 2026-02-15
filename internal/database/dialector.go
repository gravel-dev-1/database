/*
This file defines a NoOp Dialector,
It is meant to be deleted when choosing a dialector
*/

package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type noOpDialector struct{}

func (noOpDialector) Name() string                                   { return "noOpDialector" }
func (noOpDialector) Initialize(*gorm.DB) error                      { return nil }
func (noOpDialector) Migrator(*gorm.DB) gorm.Migrator                { return noOpMigrator{} }
func (noOpDialector) DataTypeOf(*schema.Field) string                { return "" }
func (noOpDialector) DefaultValueOf(*schema.Field) clause.Expression { return nil }
func (noOpDialector) BindVarTo(clause.Writer, *gorm.Statement, any)  {}
func (noOpDialector) QuoteTo(clause.Writer, string)                  {}
func (noOpDialector) Explain(string, ...any) string                  { return "" }

type noOpMigrator struct{}

func (noOpMigrator) AddColumn(any, string) error                                   { return nil }
func (noOpMigrator) AlterColumn(any, string) error                                 { return nil }
func (noOpMigrator) AutoMigrate(...any) error                                      { return nil }
func (noOpMigrator) ColumnTypes(any) ([]gorm.ColumnType, error)                    { return nil, nil }
func (noOpMigrator) CreateConstraint(any, string) error                            { return nil }
func (noOpMigrator) CreateIndex(any, string) error                                 { return nil }
func (noOpMigrator) CreateTable(...any) error                                      { return nil }
func (noOpMigrator) CreateView(string, gorm.ViewOption) error                      { return nil }
func (noOpMigrator) CurrentDatabase() string                                       { return "" }
func (noOpMigrator) DropColumn(any, string) error                                  { return nil }
func (noOpMigrator) DropConstraint(any, string) error                              { return nil }
func (noOpMigrator) DropIndex(any, string) error                                   { return nil }
func (noOpMigrator) DropTable(...any) error                                        { return nil }
func (noOpMigrator) DropView(string) error                                         { return nil }
func (noOpMigrator) FullDataTypeOf(*schema.Field) clause.Expr                      { return clause.Expr{} }
func (noOpMigrator) GetIndexes(any) ([]gorm.Index, error)                          { return nil, nil }
func (noOpMigrator) GetTables() ([]string, error)                                  { return nil, nil }
func (noOpMigrator) GetTypeAliases(string) []string                                { return nil }
func (noOpMigrator) HasColumn(any, string) bool                                    { return false }
func (noOpMigrator) HasConstraint(any, string) bool                                { return false }
func (noOpMigrator) HasIndex(any, string) bool                                     { return false }
func (noOpMigrator) HasTable(any) bool                                             { return false }
func (noOpMigrator) MigrateColumn(any, *schema.Field, gorm.ColumnType) error       { return nil }
func (noOpMigrator) MigrateColumnUnique(any, *schema.Field, gorm.ColumnType) error { return nil }
func (noOpMigrator) RenameColumn(any, string, string) error                        { return nil }
func (noOpMigrator) RenameIndex(any, string, string) error                         { return nil }
func (noOpMigrator) RenameTable(any, any) error                                    { return nil }
func (noOpMigrator) TableType(any) (gorm.TableType, error)                         { return nil, nil }
