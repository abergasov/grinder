<?php

use Phinx\Migration\AbstractMigration;

class Rights extends AbstractMigration
{
    /**
     * Change Method.
     *
     * Write your reversible migrations using this method.
     *
     * More information on writing migrations is available here:
     * http://docs.phinx.org/en/latest/migrations.html#the-abstractmigration-class
     *
     * The following commands can be used in this method and Phinx will
     * automatically reverse them when rolling back:
     *
     *    createTable
     *    renameTable
     *    addColumn
     *    addCustomColumn
     *    renameColumn
     *    addIndex
     *    addForeignKey
     *
     * Any other destructive changes will result in an error when trying to
     * rollback the migration.
     *
     * Remember to call "create()" or "update()" and NOT "save()" when working
     * with the Table class.
     */
    public function change()
    {
        $this->execute('create table users_rights
                        (
                        	ur_id int auto_increment,
                        	user_id int null,
                        	right_id int null,
                        	constraint users_rights_pk
                        		primary key (ur_id)
                        );');
        $this->execute('create index users_rights_user_id_index on users_rights (user_id);');
        $this->execute('create unique index users_rights_user_id_right_id_uindex on users_rights (user_id, right_id);');
    }
}
