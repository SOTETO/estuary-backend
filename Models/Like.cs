using System;
using System.ComponentModel.DataAnnotations.Schema;

namespace estuary_backend.Models
{
    public class Like
    {
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public int Id { get; set; }

        public string UserId { get; set; }
    }
}
