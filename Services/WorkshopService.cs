using estuary_backend.Models;
using System.Linq;
using Microsoft.EntityFrameworkCore;

namespace estuary_backend.Services
{
    public interface IWorkshopService
    {
        Workshop GetWorkshopById(int id);

        int CreateWorkshop(Workshop workshop);

        void DeleteWorkshop(int id);
    }

    public class WorkshopService : IWorkshopService
    {
        public int CreateWorkshop(Workshop workshop)
        {
            using var ctx = new EstuaryDbContext();
            ctx.Add(workshop);
            ctx.SaveChanges();
            return workshop.Id;

        }

        public Workshop GetWorkshopById(int id)
        {
            using var ctx = new EstuaryDbContext();
            var workshops = ctx.Workshops
                .Where(ws => ws.Id == id)
                .Include(ws => ws.Tags)
                .Include(ws => ws.Content)
                .Include(ws => ws.Authors)
                .FirstOrDefault();

            return workshops;
        }

        public void DeleteWorkshop(int id)
        {
            using var ctx = new EstuaryDbContext();
            var workshop = GetWorkshopById(id);
            if (workshop is not null)
            {
                ctx.Remove(workshop);
                ctx.SaveChanges();
            }
        }
    }
}
